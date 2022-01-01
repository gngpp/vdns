package vjson

import (
	"encoding/json"
	"errors"
	"fmt"
)

func PrettifyString(i interface{}) string {
	resp, _ := json.MarshalIndent(i, "", "   ")
	return string(resp)
}

func Prettify(data interface{}) ([]byte, error) {
	return json.MarshalIndent(data, "", "   ")
}

// Convert is use convert map[string]interface object to struct
func Convert(in interface{}, out interface{}) error {
	byt, _ := json.Marshal(in)
	err := json.Unmarshal(byt, out)
	return err
}

func ByteArrayConver(bytes []byte, out interface{}) error {
	return json.Unmarshal(bytes, out)
}

// Recover is use convert map[string]interface object to struct
func Recover(in interface{}) error {
	if in == nil {
		return nil
	}
	return errors.New(fmt.Sprint(in))
}

func MapToJson(data map[string]interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func StructToString(data interface{}) string {
	jsonData, err := Prettify(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

// ToMap JSON (json转map)
func ToMap(data string) map[string]interface{} {
	var jsonData map[string]interface{}
	json.Unmarshal([]byte(data), &jsonData)
	return jsonData
}

// ToIntArray
// JSONstring (json转IntList)
func ToIntArray(data string) []int {
	var tmp = make([]int, 0)
	json.Unmarshal([]byte(data), &tmp)
	return tmp
}
