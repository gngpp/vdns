package action

import (
	"vdns/lib/util/strs"
)

type RequestAction struct {
	Describe *string
	Create   *string
	Update   *string
	Delete   *string
}

func NewAliDNSAction() *RequestAction {
	return &RequestAction{
		Describe: strs.String("DescribeDomainRecords"),
		Create:   strs.String("AddDomainRecord"),
		Update:   strs.String("UpdateDomainRecord"),
		Delete:   strs.String("DeleteDomainRecord"),
	}
}

func NewDNSPodAction() *RequestAction {
	return &RequestAction{
		Describe: strs.String("DescribeRecordList"),
		Create:   strs.String("CreateRecord"),
		Update:   strs.String("ModifyRecord"),
		Delete:   strs.String("DeleteRecord"),
	}
}

func NewCloudflareAction() *RequestAction {
	return &RequestAction{
		Describe: strs.String("DescribeRecordList"),
		Create:   strs.String("CreateRecord"),
		Update:   strs.String("ModifyRecord"),
		Delete:   strs.String("DeleteRecord"),
	}
}
