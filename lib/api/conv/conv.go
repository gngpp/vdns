package conv

import (
	"vdns/lib/api/models"
	"vdns/lib/api/models/aliyun_model"
	"vdns/lib/standard/record"
)

//goland:noinspection GoRedundantConversion
func AiiyunBodyToResponse(body *aliyun_model.AliyunDescribeDomainRecordsResponseBody) *models.ApiDomainRecordResponse {
	response := &models.ApiDomainRecordResponse{}
	response.TotalCount = body.TotalCount
	response.PageSize = body.PageSize
	response.PageNumber = body.PageNumber
	aliyunRecords := body.DomainRecords.Record
	if aliyunRecords != nil {
		records := make([]*models.Record, len(aliyunRecords))
		for i, aliyunRecord := range aliyunRecords {
			record := &models.Record{
				ID:         aliyunRecord.RecordId,
				RecordType: record.Type(*aliyunRecord.Type),
				Domain:     aliyunRecord.DomainName,
				RR:         aliyunRecord.RR,
				Value:      aliyunRecord.Value,
				TTL:        aliyunRecord.TTL,
			}
			records[i] = record
		}
		response.Records = records
	}
	return response
}
