package logic

import (
	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	// 通过带缓存的 PostModel 查询文章详情
	post, err := l.svcCtx.PostModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	response := &blog.GetPostResponse{
		Id:        post.Id,
		UserId:    post.UserId,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt.Unix(),
		UpdatedAt: post.UpdatedAt.Unix(),
	}

	return response, nil
}
