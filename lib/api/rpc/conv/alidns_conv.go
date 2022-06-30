package conv

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"vdns/lib/api/errs"
	"vdns/lib/api/model"
	"vdns/lib/api/model/alidns_model"
	"vdns/lib/standard/record"
	"vdns/lib/util/iotool"
	"vdns/lib/util/vhttp"
	"vdns/lib/util/vjson"
	"vdns/lib/vlog"
)

type AliDNSResponseConvert struct {
}

//goland:noinspection GoRedundantConversion
func (_this *AliDNSResponseConvert) DescribeResponseCtxConvert(_ context.Context, resp *http.Response) (*model.DomainRecordsResponse, error) {
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
		vlog.Debugf("[AliDNSResponseConvert] json body: %v", string(bytes))
		source := new(alidns_model.DescribeDomainRecordsResponse)
		err = vjson.ByteArrayConvert(bytes, source)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		target := new(model.DomainRecordsResponse)
		target.TotalCount = source.TotalCount
		target.PageSize = source.PageSize
		target.PageNumber = source.PageNumber
		aliyunRecords := source.DomainRecords.Record
		if aliyunRecords != nil || len(aliyunRecords) > 0 {
			records := make([]*model.Record, len(aliyunRecords))
			for i, aliyunRecord := range aliyunRecords {
				if aliyunRecord != nil {
					target := &model.Record{
						ID:         aliyunRecord.RecordId,
						RecordType: record.Type(*aliyunRecord.Type),
						Domain:     aliyunRecord.DomainName,
						RR:         aliyunRecord.RR,
						Value:      aliyunRecord.Value,
						Status:     aliyunRecord.Status,
						TTL:        aliyunRecord.TTL,
					}
					records[i] = target
				}
			}
			listCount := int64(len(records))
			target.Records = records
			target.ListCount = &listCount
		}
		return target, nil
	}
	return nil, _this.badBodyHandler(body)
}

func (_this *AliDNSResponseConvert) CreateResponseCtxConvert(_ context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error) {
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
		source := new(alidns_model.CreateDomainRecordResponse)
		err = vjson.ByteArrayConvert(bytes, source)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		target := new(model.DomainRecordStatusResponse)
		return target.SetRecordId(source.RecordId).SetRequestId(source.RequestId), nil
	}
	return nil, _this.badBodyHandler(resp.Body)
}

func (_this *AliDNSResponseConvert) UpdateResponseCtxConvert(_ context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error) {
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
		source := new(alidns_model.UpdateDomainRecordResponse)
		err = vjson.ByteArrayConvert(bytes, source)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		target := new(model.DomainRecordStatusResponse)
		return target.SetRecordId(source.RecordId).SetRequestId(source.RequestId), nil
	}
	return nil, _this.badBodyHandler(resp.Body)
}

func (_this *AliDNSResponseConvert) DeleteResponseCtxConvert(_ context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error) {
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
		source := new(alidns_model.DeleteDomainRecordResponse)
		err = vjson.ByteArrayConvert(bytes, source)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		target := new(model.DomainRecordStatusResponse)
		return target.SetRecordId(source.RecordId).SetRequestId(source.RequestId), nil
	}
	return nil, _this.badBodyHandler(resp.Body)
}

//goland:noinspection GoRedundantConversion
func (_this *AliDNSResponseConvert) DescribeResponseConvert(resp *http.Response) (*model.DomainRecordsResponse, error) {
	return _this.DescribeResponseCtxConvert(nil, resp)
}

func (_this *AliDNSResponseConvert) CreateResponseConvert(resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	return _this.CreateResponseCtxConvert(nil, resp)
}

func (_this *AliDNSResponseConvert) UpdateResponseConvert(resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	return _this.UpdateResponseCtxConvert(nil, resp)
}

func (_this *AliDNSResponseConvert) DeleteResponseConvert(resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	return _this.DeleteResponseCtxConvert(nil, resp)
}

func (_this *AliDNSResponseConvert) badBodyHandler(read io.ReadCloser) error {
	bytes, err := ioutil.ReadAll(read)
	if err != nil {
		return errs.NewVdnsFromError(err)
	}
	vlog.Debugf("[AliDNSResponseConvert] bad body: %v", string(bytes))
	sdkError := new(errs.AlidnsSDKError)
	err = vjson.ByteArrayConvert(bytes, sdkError)
	if err != nil {
		return errs.NewVdnsFromError(err)
	}
	return errs.NewVdnsFromError(sdkError)
}
