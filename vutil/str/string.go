package str

import util "vdns/vutil"

//goland:noinspection ALL
func IsEmpty(value string) bool {
	return value == ""
}

//goland:noinspection ALL
func String(a string) *string {
	return &a
}

//goland:noinspection ALL
func StringValue(a *string) string {
	if a == nil {
		return ""
	}
	util.Escape("sd")
	return *a
}
