package response

import (
	"net/http"
	"reflect"

	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	body := Body{Code: 0, Message: successMessage(resp), Data: resp}
	if err != nil {
		body.Code = errorCode(err)
		body.Message = err.Error()
		body.Data = nil
	}
	httpx.OkJson(w, body)
}

func Error(w http.ResponseWriter, status int, message string) {
	httpx.WriteJson(w, status, Body{Code: errorCodeFromHTTP(status), Message: message, Data: nil})
}

func successMessage(resp interface{}) string {
	if resp == nil {
		return "操作成功"
	}
	t := reflect.TypeOf(resp)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	name := t.Name()
	switch name {
	case "CreatePostResponse":
		return "成功创建文章"
	case "CreateCategoryResponse":
		return "成功创建分类"
	case "UploadImageResponse":
		return "成功上传图片"
	case "EmptyResponse":
		return "操作成功"
	default:
		return "请求成功"
	}
}

func errorCode(err error) int {
	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.InvalidArgument:
			return 40001
		case codes.Unauthenticated:
			return 40101
		case codes.PermissionDenied:
			return 40301
		case codes.NotFound:
			return 40401
		case codes.AlreadyExists:
			return 40901
		case codes.ResourceExhausted:
			return 42901
		}
	}
	return 50001
}

func errorCodeFromHTTP(httpStatus int) int {
	switch httpStatus {
	case http.StatusBadRequest:
		return 40001
	case http.StatusUnauthorized:
		return 40101
	case http.StatusForbidden:
		return 40301
	case http.StatusNotFound:
		return 40401
	case http.StatusTooManyRequests:
		return 42901
	default:
		return 50001
	}
}
