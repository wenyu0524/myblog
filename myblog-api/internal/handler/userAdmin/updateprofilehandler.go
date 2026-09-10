package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"myblog-api/internal/logic/userAdmin"
	"myblog-api/internal/svc"
	"myblog-api/internal/types"
	"myblog-api/response"
	"net/http"
)

func UpdateProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateProfileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := userAdmin.NewUpdateProfileLogic(r.Context(), svcCtx)
		resp, err := l.UpdateProfile(&req)
		response.Response(w, resp, err)

	}
}
