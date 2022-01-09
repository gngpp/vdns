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
	"vdns/util/vjson"
)

type AlidnsResponseConvert struct {
}

//goland:noinspection GoRedundantConversion
func (_this *AlidnsResponseConvert) DescribeResponseCtxConvert(_ context.Context, resp *http.Response) (*models.DomainRecordsResponse, error) {
	if resp == nil {
		return nil, errs.NewVdnsError("*http.Response cannot been null.")
	}
	body := resp.Body
	defer body.Close()
	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		body := &alidns_model.DescribeDomainRecordsResponse{}
		err = vjson.ByteArrayConver(bytes, body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		response := &models.DomainRecordsResponse{}
		response.TotalCount = body.TotalCount
		response.PageSize = body.PageSize
		response.PageNumber = body.PageNumber
		aliyunRecords := body.DomainRecords.Record
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

func (_this *AlidnsResponseConvert) CreateResponseCtxConvert(_ context.Context, resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	if resp == nil {
		return nil, errs.NewVdnsError("*http.Response cannot been null.")
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		sourceBody := &alidns_model.CreateDomainRecordResponseBody{}
		err = vjson.ByteArrayConver(bytes, sourceBody)
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

func (_this *AlidnsResponseConvert) UpdateResponseCtxConvert(_ context.Context, resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	if resp == nil {
		return nil, errs.NewVdnsError("*http.Response cannot been null.")
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		sourceBody := &alidns_model.UpdateDomainRecordResponse{}
		err = vjson.ByteArrayConver(bytes, sourceBody)
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

func (_this *AlidnsResponseConvert) DeleteResponseCtxConvert(_ context.Context, resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	if resp == nil {
		return nil, errs.NewVdnsError("*http.Response cannot been null.")
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		sourceBody := &alidns_model.DeleteDomainRecordResponse{}
		err = vjson.ByteArrayConver(bytes, sourceBody)
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
func (_this *AlidnsResponseConvert) DescribeResponseConvert(resp *http.Response) (*models.DomainRecordsResponse, error) {
	return _this.DescribeResponseCtxConvert(nil, resp)
}

func (_this *AlidnsResponseConvert) CreateResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	return _this.CreateResponseCtxConvert(nil, resp)
}

func (_this *AlidnsResponseConvert) UpdateResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	return _this.UpdateResponseCtxConvert(nil, resp)
}

func (_this *AlidnsResponseConvert) DeleteResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	return _this.DeleteResponseCtxConvert(nil, resp)
}

func (_this *AlidnsResponseConvert) badBodyHandler(read io.ReadCloser) error {
	bytes, err := ioutil.ReadAll(read)
	if err != nil {
		return errs.NewApiErrorFromError(err)
	}
	sdkError := &errs.AlidnsSDKError{}
	err = vjson.ByteArrayConver(bytes, sdkError)
	if err != nil {
		return errs.NewApiErrorFromError(err)
	}
	return errs.NewApiErrorFromError(sdkError)
}
