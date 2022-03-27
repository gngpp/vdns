package conv

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"vdns/lib/api/errs"
	"vdns/lib/api/models"
	"vdns/lib/api/models/alidns_model"
	"vdns/lib/standard/record"
	"vdns/lib/util/vjson"
	"vdns/lib/vlog"
)

type AliDNSResponseConvert struct {
}

//goland:noinspection GoRedundantConversion
func (_this *AliDNSResponseConvert) DescribeResponseCtxConvert(_ context.Context, resp *http.Response) (*models.DomainRecordsResponse, error) {
	if resp == nil {
		return nil, errs.NewVdnsError("*http.Response cannot been null.")
	}
	body := resp.Body
	defer _this.closeBody(body)
	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		sourceResponse := new(alidns_model.DescribeDomainRecordsResponse)
		err = vjson.ByteArrayConvert(bytes, sourceResponse)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		response := new(models.DomainRecordsResponse)
		response.TotalCount = sourceResponse.TotalCount
		response.PageSize = sourceResponse.PageSize
		response.PageNumber = sourceResponse.PageNumber
		aliyunRecords := sourceResponse.DomainRecords.Record
		if aliyunRecords != nil || len(aliyunRecords) > 0 {
			records := make([]*models.Record, len(aliyunRecords))
			for i, aliyunRecord := range aliyunRecords {
				if aliyunRecord != nil {
					target := &models.Record{
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
			response.Records = records
			response.ListCount = &listCount
		}
		return response, nil
	} else {
		return nil, _this.badBodyHandler(body)
	}
}

func (_this *AliDNSResponseConvert) CreateResponseCtxConvert(_ context.Context, resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	if resp == nil {
		return nil, errs.NewVdnsError("*http.Response cannot been null.")
	}
	body := resp.Body
	defer _this.closeBody(body)
	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		sourceResponse := new(alidns_model.CreateDomainRecordResponse)
		err = vjson.ByteArrayConvert(bytes, sourceResponse)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		response := &models.DomainRecordStatusResponse{
			RecordId:  sourceResponse.RecordId,
			RequestId: sourceResponse.RequestId,
		}
		return response, nil
	} else {
		return nil, _this.badBodyHandler(resp.Body)
	}
}

func (_this *AliDNSResponseConvert) UpdateResponseCtxConvert(_ context.Context, resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	if resp == nil {
		return nil, errs.NewVdnsError("*http.Response cannot been null.")
	}
	body := resp.Body
	defer _this.closeBody(body)
	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		sourceResponse := new(alidns_model.UpdateDomainRecordResponse)
		err = vjson.ByteArrayConvert(bytes, sourceResponse)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		response := &models.DomainRecordStatusResponse{
			RecordId:  sourceResponse.RecordId,
			RequestId: sourceResponse.RequestId,
		}
		return response, nil
	} else {
		return nil, _this.badBodyHandler(resp.Body)
	}
}

func (_this *AliDNSResponseConvert) DeleteResponseCtxConvert(_ context.Context, resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	if resp == nil {
		return nil, errs.NewVdnsError("*http.Response cannot been null.")
	}
	body := resp.Body
	defer _this.closeBody(body)
	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		sourceBody := new(alidns_model.DeleteDomainRecordResponse)
		err = vjson.ByteArrayConvert(bytes, sourceBody)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		response := &models.DomainRecordStatusResponse{
			RecordId:  sourceBody.RecordId,
			RequestId: sourceBody.RequestId,
		}
		return response, nil
	} else {
		return nil, _this.badBodyHandler(resp.Body)
	}
}

//goland:noinspection GoRedundantConversion
func (_this *AliDNSResponseConvert) DescribeResponseConvert(resp *http.Response) (*models.DomainRecordsResponse, error) {
	return _this.DescribeResponseCtxConvert(nil, resp)
}

func (_this *AliDNSResponseConvert) CreateResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	return _this.CreateResponseCtxConvert(nil, resp)
}

func (_this *AliDNSResponseConvert) UpdateResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	return _this.UpdateResponseCtxConvert(nil, resp)
}

func (_this *AliDNSResponseConvert) DeleteResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	return _this.DeleteResponseCtxConvert(nil, resp)
}

func (_this *AliDNSResponseConvert) badBodyHandler(read io.ReadCloser) error {
	bytes, err := ioutil.ReadAll(read)
	if err != nil {
		return errs.NewApiErrorFromError(err)
	}
	sdkError := new(errs.AlidnsSDKError)
	err = vjson.ByteArrayConvert(bytes, sdkError)
	if err != nil {
		return errs.NewApiErrorFromError(err)
	}
	return errs.NewApiErrorFromError(sdkError)
}

func (_this AliDNSResponseConvert) closeBody(body io.ReadCloser) {
	if body == nil {
		return
	}
	err := body.Close()
	if err != nil {
		vlog.Error(err)
	}
}
