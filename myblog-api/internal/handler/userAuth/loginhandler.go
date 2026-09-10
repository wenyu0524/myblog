package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"myblog-api/internal/logic/userAuth"
	"myblog-api/internal/svc"
	"myblog-api/internal/types"
	"myblog-api/response"
	"net/http"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := userAuth.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err)

	}
}
