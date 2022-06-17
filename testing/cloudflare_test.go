package testing

import (
	"fmt"
	"testing"
	"vdns/lib/api"
	"vdns/lib/api/model"
	"vdns/lib/auth"
	"vdns/lib/standard/record"
	"vdns/lib/util/vjson"
)

func TestCFDescribe(t *testing.T) {
	//vlog.SetLevel(vlog.Level.DEBUG)
	credential := auth.NewTokenCredential("shww4JpWY1Ilp43DHDMwY8ja_aoPs-RSJwmTcobi")
	provider := api.NewCloudflareProvider(credential)
	err := provider.Support(record.A)
	if err != nil {
		fmt.Println(err)
	} else {
		req := model.NewDescribeDomainRecordsRequest()
		req.SetDomain("innas.work")
		//req.SetRecordType(record.A)
		req.SetPageSize(5)
		req.SetPageNumber(1)
		// 不支持模糊匹配
		req.SetValueKeyWord("101.35.6.187")
		// cloudflare 不支持
		//req.SetRRKeyWord("")
		records, err := provider.DescribeRecords(req)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(vjson.PrettifyString(records))
	}
}

func TestCFDelete(t *testing.T) {
	//vlog.SetLevel(vlog.Level.DEBUG)
	credential := auth.NewTokenCredential("shww4JpWY1Ilp43DHDMwY8ja_aoPs-RSJwmTcobi")
	provider := api.NewCloudflareProvider(credential)
	err := provider.Support(record.A)
	if err != nil {
		fmt.Println(err)
	} else {
		req := model.NewDeleteDomainRecordRequest()
		req.SetDomain("innas.work")
		req.SetID("856b72a4d916c9e44f545f9480123ca9")
		// cloudflare 不支持
		//req.SetRRKeyWord("")
		records, err := provider.DeleteRecord(req)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(vjson.PrettifyString(records))
	}
}

func TestCFCreate(t *testing.T) {
	//vlog.SetLevel(vlog.Level.DEBUG)
	credential := auth.NewTokenCredential("shww4JpWY1Ilp43DHDMwY8ja_aoPs-RSJwmTcobi")
	provider := api.NewCloudflareProvider(credential)
	err := provider.Support(record.A)
	if err != nil {
		fmt.Println(err)
	} else {
		req := model.NewCreateDomainRecordRequest()
		req.SetDomain("17.innas.work")
		req.SetRecordType(record.A)
		req.SetValue("127.0.0.1")
		// cloudflare 不支持
		//req.SetRRKeyWord("")
		records, err := provider.CreateRecord(req)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(vjson.PrettifyString(records))
	}
}
