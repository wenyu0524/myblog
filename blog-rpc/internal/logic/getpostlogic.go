package logic

import (
	"context"
	"errors"

	"blog-rpc/internal/model"
	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLogic {
	return &GetPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetPost 根据文章 ID 查询文章详情，将时间戳转为 Unix 秒返回
func (l *GetPostLogic) GetPost(in *blog.GetPostRequest) (*blog.GetPostResponse, error) {
	// FindOne 由 goctl 生成的 CachedConn 实现，按主键自动读取/回填 Redis。
	post, err := l.svcCtx.PostsModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	return &blog.GetPostResponse{
		Id:        post.Id,
		UserId:    post.UserId,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt.Unix(),
		UpdatedAt: post.UpdatedAt.Unix(),
	}, nil
}
