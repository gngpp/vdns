package msg

// api message
//goland:noinspection ALL
const (
	RECORD_TYPE_NOT_SUPPORT  = "Record type is not supported."
	RECORD_VALUE_NOT_SUPPORT = "Record value cannot been nil."
	RECORD_ID_NOT_SUPPORT    = "Record id cannot been nil."
	DOMAIN_NOT_SPECIFICATION = "the domain name does not meet the specification."
	DESCRIBE_REQUEST_NOT_NIL = "DescribeDomainRecordsRequest cannot been nil."
	CREATE_REQUEST_NOT_NIL   = "CreateDomainRecordRequest cannot been nil."
	UPDATE_REQUEST_NOT_NIL   = "UpdateDomainRecordRequest cannot been nil."
	DELETE_REQUEST_NOT_NIL   = "DeleteDomainRecordRequest cannot been nil."
)

//goland:noinspection ALL
const (
	SYSTEM_CREDENTIAL_NOT_NIL         = "credential can not been nil."
	SYSTEM_SIGNATURE_COMPOSER_NOT_NIL = "signature composer cannot been nil."
)
