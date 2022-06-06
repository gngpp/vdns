package conv

import (
	"context"
	"io/ioutil"
	"net/http"
	"vdns/lib/api/errs"
	"vdns/lib/api/model"
	"vdns/lib/api/model/dnspod_model"
	"vdns/lib/api/parameter"
	"vdns/lib/standard/record"
	"vdns/lib/util/convert"
	"vdns/lib/util/iotool"
	"vdns/lib/util/strs"
	"vdns/lib/util/vhttp"
	"vdns/lib/util/vjson"
)

type DNSPodResponseConvert struct {
}

//goland:noinspection GoRedundantConversion
func (_this *DNSPodResponseConvert) DescribeResponseCtxConvert(ctx context.Context, resp *http.Response) (*model.DomainRecordsResponse, error) {
	if resp == nil {
		return nil, errs.NewVdnsError("*http.Response cannot been null.")
	}
	ctxDescribeRequest := new(model.DescribeDomainRecordsRequest)
	if ctx != nil {
		request := ctx.Value(parameter.DnspocParameterContextDescribeKey)
		err := vjson.Convert(request, ctxDescribeRequest)
		if err != nil {
			return nil, err
		}
	}
	body := resp.Body
	defer iotool.ReadCloser(body)
	if vhttp.IsOK(resp) {
		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		b := new(dnspod_model.DescribeRecordListResponse)
		err = vjson.ByteArrayConvert(bytes, b)
		if err != nil {
			return nil, errs.NewVdnsFromError(err)
		}
		sourceResponse := b.Response
		if sourceResponse != nil {
			if sourceResponse.Error != nil {
				return nil, _this.errorBodyHandler(sourceResponse.Error, sourceResponse.RequestId)
			}
			dnspodRecords := sourceResponse.RecordList
			if dnspodRecords != nil || len(dnspodRecords) > 0 {
				records := make([]*model.Record, len(dnspodRecords))
				for i, dnspodRecord := range dnspodRecords {
					if dnspodRecord != nil {
						target := &model.Record{
							ID:         convert.AsString(dnspodRecord.RecordId),
							RecordType: record.Type(*dnspodRecord.Type),
							Domain:     ctxDescribeRequest.Domain,
							RR:         dnspodRecord.Name,
							Value:      dnspodRecord.Value,
							Status:     dnspodRecord.Status,
							TTL:        dnspodRecord.TTL,
						}
						records[i] = target
					}
				}
				var pageSize int64 = 100
				if ctxDescribeRequest.PageSize != nil {
					pageSize = *ctxDescribeRequest.PageSize
				}
				listCount := int64(len(records))
				response := &model.DomainRecordsResponse{
					TotalCount: sourceResponse.RecordCountInfo.TotalCount,
					PageSize:   &pageSize,
					PageNumber: ctxDescribeRequest.PageNumber,
					Records:    records,
					ListCount:  &listCount,
				}
				return response, nil
			}
		}
		return &model.DomainRecordsResponse{}, nil
	}
	return &model.DomainRecordsResponse{}, nil
}

func (_this *DNSPodResponseConvert) CreateResponseCtxConvert(_ context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error) {
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
		c := new(dnspod_model.CreateRecordResponse)
		err = vjson.ByteArrayConvert(bytes, c)
		if err != nil {
			return nil, err
		}
		sourceResponse := c.Response
		if sourceResponse.Error != nil {
			return nil, _this.errorBodyHandler(sourceResponse.Error, sourceResponse.RequestId)
		}
		response := &model.DomainRecordStatusResponse{
			RecordId:  convert.AsString(sourceResponse.RecordId),
			RequestId: sourceResponse.RequestId,
		}
		return response, nil
	} else {
		return nil, errs.NewVdnsError("dnspod bad response.")
	}
}

func (_this *DNSPodResponseConvert) UpdateResponseCtxConvert(_ context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error) {
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
		c := new(dnspod_model.ModifyRecordResponse)
		err = vjson.ByteArrayConvert(bytes, c)
		if err != nil {
			return nil, err
		}
		sourceResponse := c.Response
		if sourceResponse.Error != nil {
			return nil, _this.errorBodyHandler(sourceResponse.Error, sourceResponse.RequestId)
		}
		response := &model.DomainRecordStatusResponse{
			RecordId:  convert.AsString(sourceResponse.RecordId),
			RequestId: sourceResponse.RequestId,
		}
		return response, nil
	} else {
		return nil, errs.NewVdnsError("dnspod bad response.")
	}
}

func (_this *DNSPodResponseConvert) DeleteResponseCtxConvert(_ context.Context, resp *http.Response) (*model.DomainRecordStatusResponse, error) {
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
		c := new(dnspod_model.DeleteRecordResponse)
		err = vjson.ByteArrayConvert(bytes, c)
		if err != nil {
			return nil, err
		}
		sourceResponse := c.Response
		if sourceResponse.Error != nil {
			return nil, _this.errorBodyHandler(sourceResponse.Error, sourceResponse.RequestId)
		}
		response := &model.DomainRecordStatusResponse{
			RecordId:  convert.AsString("none"),
			RequestId: sourceResponse.RequestId,
		}
		return response, nil
	} else {
		return nil, errs.NewVdnsError("dnspod bad response.")
	}
}

func (_this *DNSPodResponseConvert) DescribeResponseConvert(resp *http.Response) (*model.DomainRecordsResponse, error) {
	return _this.DescribeResponseCtxConvert(nil, resp)
}

func (_this *DNSPodResponseConvert) CreateResponseConvert(resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	return _this.CreateResponseCtxConvert(nil, resp)
}

func (_this *DNSPodResponseConvert) UpdateResponseConvert(resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	return _this.UpdateResponseCtxConvert(nil, resp)
}

func (_this *DNSPodResponseConvert) DeleteResponseConvert(resp *http.Response) (*model.DomainRecordStatusResponse, error) {
	return _this.DeleteResponseCtxConvert(nil, resp)
}

func (_this *DNSPodResponseConvert) errorBodyHandler(e *dnspod_model.Error, requestId *string) error {
	return errs.NewVdnsFromError(errs.NewTencentCloudSDKError(strs.StringValue(e.Code),
		strs.StringValue(e.Message), strs.StringValue(requestId)))
}
