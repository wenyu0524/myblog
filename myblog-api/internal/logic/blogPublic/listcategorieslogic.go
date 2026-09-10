// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package blogPublic

import (
	"blog-rpc/blogclient"
	"context"

	"myblog-api/internal/svc"
	"myblog-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCategoriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCategoriesLogic {
	return &ListCategoriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCategoriesLogic) ListCategories() (resp *types.ListCategoriesResponse, err error) {
	r, err := l.svcCtx.BlogRpc.ListCategories(l.ctx, &blogclient.Empty{})
	if err != nil {
		return nil, err
	}
	resp = &types.ListCategoriesResponse{Categories: make([]types.CategoryItem, 0, len(r.Categories))}
	for _, c := range r.Categories {
		resp.Categories = append(resp.Categories, types.CategoryItem{
			Id:        c.Id,
			Name:      c.Name,
			Slug:      c.Slug,
			Sort:      c.Sort,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}
	return resp, nil
}
