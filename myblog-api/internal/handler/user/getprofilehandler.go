package handler

import (
	"myblog-api/internal/logic/user"
	"myblog-api/internal/svc"
	"myblog-api/response"
	"net/http"
)

func GetProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := user.NewGetProfileLogic(r.Context(), svcCtx)
		resp, err := l.GetProfile()
		response.Response(w, resp, err)

	}
}
