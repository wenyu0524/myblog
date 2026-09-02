package logic

import (
	"context"

	"blog-rpc/internal/model"
	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreatePost 创建文章，将文章数据写入 blog_db
func (l *CreatePostLogic) CreatePost(in *blog.CreatePostRequest) (*blog.CreatePostResponse, error) {
	id, err := l.svcCtx.PostsModel.InsertReturningID(l.ctx, &model.Posts{
		UserId:  in.UserId,
		Title:   in.Title,
		Content: in.Content,
	})
	if err != nil {
		return nil, err
	}

	return &blog.CreatePostResponse{Id: id}, nil
}
