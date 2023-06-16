package helper

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(code int, message string) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["status"] = http.StatusText(code)
	resp["code"] = code
	resp["message"] = message
	return resp
}

func AccessToken(code int, access_token string) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["status"] = http.StatusText(code)
	resp["code"] = code
	resp["access_token"] = access_token
	return resp
}

func Append(json1, json2 []byte) ([]byte, error) {
	var jsonArray1 []json.RawMessage
	err := json.Unmarshal(json1, &jsonArray1)
	if err != nil {
		return nil, err
	}
	var jsonArray2 []json.RawMessage
	err = json.Unmarshal(json2, &jsonArray2)
	if err != nil {
		return nil, err
	}
	jsonArray := append(jsonArray1, jsonArray2...)
	result, err := json.Marshal(jsonArray)
	if err != nil {
		return nil, err
	}
	return result, nil
}
