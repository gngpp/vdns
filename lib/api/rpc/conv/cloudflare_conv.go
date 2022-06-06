package conv

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"vdns/lib/api/errs"
	"vdns/lib/api/model"
	"vdns/lib/api/model/cloudflare_model"
	record2 "vdns/lib/standard/record"
	"vdns/lib/util/convert"
	"vdns/lib/util/iotool"
	"vdns/lib/util/strs"
	"vdns/lib/util/vhttp"
	"vdns/lib/util/vjson"
)

type CloudflareResponseConvert struct {
}

func (_this *CloudflareResponseConvert) DescribeResponseConvert(resp *http.Response) (*model.DomainRecordsResponse, error) {
	return _this.DescribeResponseCtxConvert(nil, resp)
}

func (_this *CloudflareResponseConvert) CreateResponseConvert(resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	return _this.CreateResponseCtxConvert(nil, resp)
}

func (_this *CloudflareResponseConvert) UpdateResponseConvert(resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	return _this.UpdateResponseCtxConvert(nil, resp)
}

func (_this *CloudflareResponseConvert) DeleteResponseConvert(resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	return _this.DeleteResponseCtxConvert(nil, resp)
}

func (_this *CloudflareResponseConvert) DescribeResponseCtxConvert(_ context.Context, resp *http.Response) (*model.DomainRecordsResponse, error) {
	if resp == nil {
		return nil, errs.NewVdnsError("*http.Response cannot been null.")
	}
	body := resp.Body
	defer iotool.ReadCloser(body)
	if vhttp.IsOK(resp) {
		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		source := &cloudflare_model.DescribeRecordResponse{}
		err = vjson.ByteArrayConvert(bytes, source)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		if len(source.Errors) != 0 {
			return nil, errs.NewCloudFlareSDKError(vjson.PrettifyString(source.Errors))
		}
		if source.Success && len(source.Result) > 0 {
			target := &model.DomainRecordsResponse{}
			target.Records = make([]*model.Record, len(source.Result))
			target.PageSize = source.ResultInfo.PerPage
			target.TotalCount = source.ResultInfo.TotalCount
			target.ListCount = source.ResultInfo.Count
			target.PageNumber = source.ResultInfo.Page
			for index, result := range source.Result {
				domain, err := vhttp.CheckExtractDomain(*result.Name)
				if err != nil {
					return nil, errs.NewVdnsFromError(err)
				}
				record := &model.Record{}
				record.ID = result.ID
				record.Domain = result.ZoneName
				record.RR = &domain.SubDomain
				record.Value = result.Content
				record.RecordType = record2.Type(*result.Type)
				record.TTL = result.TTL
				record.Status = strs.String(strs.Concat("proxied=", convert.AsStringValue(result.Proxied)))
				target.Records[index] = record
			}
			return target, nil
		}
	}
	return nil, _this.errorHandler(body)
}

func (_this *CloudflareResponseConvert) CreateResponseCtxConvert(_ context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	if resp == nil {
		return nil, errs.NewVdnsError("*http.Response cannot been null.")
	}
	body := resp.Body
	defer iotool.ReadCloser(body)
	if vhttp.IsOK(resp) {
		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		source := new(cloudflare_model.CreateRecordResponse)
		err = vjson.ByteArrayConvert(bytes, source)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		if len(source.Errors) > 0 {
			return nil, errs.NewVdnsError(vjson.PrettifyString(source.Errors))
		}
		target := new(model.DomainRecordStatusResponse)
		return target.SetRecordId(source.Result.ID).SetRequestId(strs.String("none")), err
	}
	return nil, _this.errorHandler(body)
}

func (_this *CloudflareResponseConvert) UpdateResponseCtxConvert(_ context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	if resp == nil {
		return nil, errs.NewVdnsError("*http.Response cannot been null.")
	}
	body := resp.Body
	defer iotool.ReadCloser(body)
	if vhttp.IsOK(resp) {
		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		source := new(cloudflare_model.UpdateRecordResponse)
		err = vjson.ByteArrayConvert(bytes, source)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		if len(source.Errors) > 0 || !source.Success {
			return nil, errs.NewVdnsError(vjson.PrettifyString(source.Errors))
		}
		target := new(model.DomainRecordStatusResponse)
		return target.SetRecordId(source.Result.ID).SetRequestId(strs.String("none")), err
	}
	return nil, _this.errorHandler(body)
}

func (_this *CloudflareResponseConvert) DeleteResponseCtxConvert(_ context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	if resp == nil {
		return nil, errs.NewVdnsError("*http.Response cannot been null.")
	}
	body := resp.Body
	defer iotool.ReadCloser(body)
	if vhttp.IsOK(resp) {
		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		source := new(cloudflare_model.DeleteRecordResponse)
		err = vjson.ByteArrayConvert(bytes, source)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		target := new(model.DomainRecordStatusResponse)
		return target.SetRecordId(source.Result.ID).SetRequestId(strs.String("none")), err
	}
	return nil, _this.errorHandler(body)
}

func (_this *CloudflareResponseConvert) errorHandler(body io.ReadCloser) error {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	c := new(cloudflare_model.Response)
	err = vjson.ByteArrayConvert(bytes, c)
	if err != nil {
		return errs.NewVdnsFromError(err)
	}
	return errs.NewVdnsError(vjson.PrettifyString(c.Errors))
}
