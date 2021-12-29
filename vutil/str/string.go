package str

import (
	"strings"
)

//goland:noinspection ALL
func IsEmpty(value string) bool {
	return value == ""
}

func NotEmpty(value string) bool {
	return !IsEmpty(value)
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
