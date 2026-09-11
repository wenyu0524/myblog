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

type DeleteCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(in *blog.DeleteCategoryRequest) (*blog.EmptyResponse, error) {
	if in == nil || in.UserId <= 0 || in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "用户ID和分类ID不能为空")
	}
	err := l.svcCtx.CategoriesModel.Delete(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "分类不存在")
		}
		return nil, status.Error(codes.FailedPrecondition, "分类仍被文章引用或无法删除")
	}
	return &blog.EmptyResponse{}, nil
}
