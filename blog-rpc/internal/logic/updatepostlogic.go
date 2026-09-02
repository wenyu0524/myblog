package logic

import (
	"context"
	"errors"
	"time"

	"blog-rpc/internal/model"
	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdatePost 更新文章，通过 user_id 校验所有权，受影响行数为 0 表示无权操作
func (l *UpdatePostLogic) UpdatePost(in *blog.UpdatePostRequest) (*blog.UpdatePostResponse, error) {
	affected, err := l.svcCtx.PostModel.Update(l.ctx, &model.Post{
		Id:        in.Id,
		UserId:    in.UserId,
		Title:     in.Title,
		Content:   in.Content,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	// 受影响行数为 0：文章不存在或当前用户非作者
	if affected == 0 {
		return nil, errors.New("post not found or you don't have permission to update")
	}

	return &blog.UpdatePostResponse{}, nil
}
