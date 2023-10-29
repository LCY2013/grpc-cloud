package ack

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func ToSuccessResponse(date string) string {
	var dataAny any
	err := json.Unmarshal([]byte(date), &dataAny)
	if err != nil {
		dataAny = date
	}
	marshal, err := json.Marshal(Response{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: dataAny,
	})
	if err != nil {
		return ""
	}
	return string(marshal)
}

func ToFailResponse(date string) string {
	marshal, err := json.Marshal(Response{
		Code: http.StatusOK,
		Msg:  date,
	})
	if err != nil {
		return ""
	}
	return string(marshal)
}
