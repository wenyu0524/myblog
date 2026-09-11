package logic

import (
	"blog-rpc/internal/model"
	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCategoryLogic) CreateCategory(in *blog.CreateCategoryRequest) (*blog.CreateCategoryResponse, error) {
	if in == nil || in.UserId <= 0 || in.Name == "" || in.Slug == "" {
		return nil, status.Error(codes.InvalidArgument, "用户、分类名称和别名不能为空")
	}
	id, err := l.svcCtx.CategoriesModel.Create(l.ctx, &model.Categories{Name: in.Name, Slug: in.Slug, Sort: in.Sort})
	if err != nil {
		return nil, status.Error(codes.Internal, "分类创建失败，请稍后重试")
	}
	return &blog.CreateCategoryResponse{Id: id}, nil
}
