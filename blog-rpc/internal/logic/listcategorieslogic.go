package logic

import (
	"context"

	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ListCategoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCategoriesLogic {
	return &ListCategoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListCategoriesLogic) ListCategories(in *blog.Empty) (*blog.ListCategoriesResponse, error) {
	rows, err := l.svcCtx.CategoriesModel.ListPublic(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "分类查询失败，请稍后重试")
	}
	resp := &blog.ListCategoriesResponse{Categories: make([]*blog.CategoryItem, 0, len(rows))}
	for _, c := range rows {
		resp.Categories = append(resp.Categories, &blog.CategoryItem{Id: c.Id, Name: c.Name, Slug: c.Slug, Sort: c.Sort, CreatedAt: c.CreatedAt.Unix(), UpdatedAt: c.UpdatedAt.Unix()})
	}
	return resp, nil
}
