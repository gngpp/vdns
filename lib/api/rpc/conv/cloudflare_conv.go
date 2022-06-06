package conv

import (
	"context"
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

func (_this *CloudflareResponseConvert) DescribeResponseCtxConvert(ctx context.Context, resp *http.Response) (*model.DomainRecordsResponse, error) {
	if vhttp.IsOK(resp) {
		body := resp.Body
		defer iotool.ReadCloser(body)
		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		r := &cloudflare_model.DescribeRecordResponse{}
		err = vjson.ByteArrayConvert(bytes, r)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		if len(r.Errors) != 0 {
			return nil, errs.NewCloudFlareSDKError(vjson.PrettifyString(r.Errors))
		}
		if r.Success && len(r.Result) > 0 {
			target := &model.DomainRecordsResponse{}
			target.PageSize = r.ResultInfo.PerPage
			target.TotalCount = r.ResultInfo.TotalCount
			target.ListCount = r.ResultInfo.Count
			target.PageNumber = r.ResultInfo.Page
			for _, result := range r.Result {
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
				target.Records = append(target.Records, record)
			}
			return target, nil
		}
	}
	return &model.DomainRecordsResponse{}, nil
}

func (_this *CloudflareResponseConvert) CreateResponseCtxConvert(ctx context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	return &model.DomainRecordStatusResponse{}, nil
}

func (_this *CloudflareResponseConvert) UpdateResponseCtxConvert(ctx context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	return &model.DomainRecordStatusResponse{}, nil
}

func (_this *CloudflareResponseConvert) DeleteResponseCtxConvert(ctx context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	return &model.DomainRecordStatusResponse{}, nil
}
