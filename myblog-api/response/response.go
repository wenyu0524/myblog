package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	body := Body{Code: 0, Message: "OK", Data: resp}
	if err != nil {
		body.Code = -1
		body.Message = err.Error()
		body.Data = nil
	}
	httpx.OkJson(w, body)
}

func Error(w http.ResponseWriter, status int, message string) {
	httpx.WriteJson(w, status, Body{Code: -1, Message: message, Data: nil})
}
