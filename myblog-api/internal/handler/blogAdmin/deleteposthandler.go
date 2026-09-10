package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"myblog-api/internal/logic/blogAdmin"
	"myblog-api/internal/svc"
	"myblog-api/internal/types"
	"myblog-api/response"
	"net/http"
)

func DeletePostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeletePostRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := blogAdmin.NewDeletePostLogic(r.Context(), svcCtx)
		resp, err := l.DeletePost(&req)
		response.Response(w, resp, err)

	}
}
