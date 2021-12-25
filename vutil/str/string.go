package str

import (
	"strings"
	util "vdns/vutil"
)

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

func Concat(strs ...string) string {
	builder := strings.Builder{}
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}

func ToBytes(str string) []byte {
	return []byte(str)
}

func ToString(bytes []byte) string {
	return string(bytes)
}
