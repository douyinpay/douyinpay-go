package utils

import (
	"encoding/json"
	"log"
)

/**
json转为string
*/
func Json2Str(obj interface{}) string {
	jsonByte, err := json.MarshalIndent(obj, "", "	")
	if err != nil {
		log.Fatal("json format error!")
		return ""
	}

	return string(jsonByte)
}
