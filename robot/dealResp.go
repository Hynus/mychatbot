package robot

import (
	"bytes"
	"encoding/json"
)

func ParseRespJson(respBytes []byte) (int, map[string]interface{}) {
	var (
		jsonTag int
		ret     map[string]interface{}
	)
	if bytes.Contains(respBytes, []byte("100000")) {
		jsonTag = 1
	} else if bytes.Contains(respBytes, []byte("200000")) {
		jsonTag = 2
	} else if bytes.Contains(respBytes, []byte("302000")) {
		jsonTag = 3
	} else if bytes.Contains(respBytes, []byte("308000")) {
		jsonTag = 4
	}
	json.Unmarshal(respBytes, &ret)
	return jsonTag, ret
}
