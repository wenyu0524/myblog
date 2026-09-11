package logic

import (
	"blog-rpc/internal/model"
	"context"
	"errors"

	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *blog.UpdateCategoryRequest) (*blog.EmptyResponse, error) {
	if in == nil || in.UserId <= 0 || in.Id <= 0 || in.Name == "" || in.Slug == "" {
		return nil, status.Error(codes.InvalidArgument, "用户、分类ID、名称和别名不能为空")
	}
	err := l.svcCtx.CategoriesModel.UpdateAdmin(l.ctx, &model.Categories{Id: in.Id, Name: in.Name, Slug: in.Slug, Sort: in.Sort})
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "分类不存在")
		}
		return nil, status.Error(codes.Internal, "分类更新失败，请稍后重试")
	}
	return &blog.EmptyResponse{}, nil
}
