package handler

import (
	"myblog-api/internal/logic/userAdmin"
	"myblog-api/internal/svc"
	"myblog-api/response"
	"net/http"
)

func GetMeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := userAdmin.NewGetMeLogic(r.Context(), svcCtx)
		resp, err := l.GetMe()
		response.Response(w, resp, err)

	}
}
