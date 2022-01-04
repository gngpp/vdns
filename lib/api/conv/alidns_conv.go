package conv

import (
	"io"
	"io/ioutil"
	"net/http"
	"vdns/lib/api/errs"
	"vdns/lib/api/models"
	"vdns/lib/api/models/alidns_model"
	"vdns/lib/standard/record"
	"vdns/vutil/vjson"
)

type AlidnsDomainRecordResponseConvert struct{}

//goland:noinspection GoRedundantConversion
func (_this *AlidnsDomainRecordResponseConvert) DescribeResponseConvert(resp *http.Response) (*models.DomainRecordsResponse, error) {
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		body := &alidns_model.DescribeDomainRecordsResponse{}
		err = vjson.ByteArrayConver(bytes, body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		resp := &models.DomainRecordsResponse{}
		resp.TotalCount = body.TotalCount
		resp.PageSize = body.PageSize
		resp.PageNumber = body.PageNumber
		aliyunRecords := body.DomainRecords.Record
		if aliyunRecords != nil {
			records := make([]*models.Record, len(aliyunRecords))
			for i, aliyunRecord := range aliyunRecords {
				if aliyunRecord != nil {
					target := &models.Record{
						ID:         aliyunRecord.RecordId,
						RecordType: record.Type(*aliyunRecord.Type),
						Domain:     aliyunRecord.DomainName,
						RR:         aliyunRecord.RR,
						Value:      aliyunRecord.Value,
						TTL:        aliyunRecord.TTL,
					}
					records[i] = target
				}
			}
			resp.Records = records
		}
		return resp, nil
	} else {
		return nil, _this.badBodyHandler(resp.Body)
	}
}

func (_this *AlidnsDomainRecordResponseConvert) CreateResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		body := &alidns_model.CreateDomainRecordResponseBody{}
		err = vjson.ByteArrayConver(bytes, body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		response := &models.DomainRecordStatusResponse{
			RecordId:  body.RecordId,
			RequestId: body.RequestId,
		}
		return response, nil
	} else {
		return nil, _this.badBodyHandler(resp.Body)
	}
}

func (_this *AlidnsDomainRecordResponseConvert) UpdateResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		body := &alidns_model.UpdateDomainRecordResponse{}
		err = vjson.ByteArrayConver(bytes, body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		response := &models.DomainRecordStatusResponse{
			RecordId:  body.RecordId,
			RequestId: body.RequestId,
		}
		return response, nil
	} else {
		return nil, _this.badBodyHandler(resp.Body)
	}
}

func (_this *AlidnsDomainRecordResponseConvert) DeleteResponseConvert(resp *http.Response) (*models.DomainRecordStatusResponse, error) {
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		body := &alidns_model.DeleteDomainRecordResponse{}
		err = vjson.ByteArrayConver(bytes, body)
		if err != nil {
			return nil, errs.NewApiErrorFromError(err)
		}
		response := &models.DomainRecordStatusResponse{
			RecordId:  body.RecordId,
			RequestId: body.RequestId,
		}
		return response, nil
	} else {
		return nil, _this.badBodyHandler(resp.Body)
	}
}

func (_this *AlidnsDomainRecordResponseConvert) badBodyHandler(read io.ReadCloser) error {
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
