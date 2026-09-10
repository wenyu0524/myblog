package handler

import (
	"myblog-api/internal/logic/blogPublic"
	"myblog-api/internal/svc"
	"myblog-api/response"
	"net/http"
)

func ListCategoriesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := blogPublic.NewListCategoriesLogic(r.Context(), svcCtx)
		resp, err := l.ListCategories()
		response.Response(w, resp, err)

	}
}
