// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package blogAdmin

import (
	"blog-rpc/blogclient"
	"context"
	"errors"
	userAdmin "myblog-api/internal/logic/userAdmin"

	"myblog-api/internal/svc"
	"myblog-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLogic) CreatePost(req *types.CreatePostRequest) (resp *types.CreatePostResponse, err error) {
	if req == nil {
		return nil, errors.New("请求不能为空")
	}
	userId, err := userAdmin.UserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	r, err := l.svcCtx.BlogRpc.CreatePost(l.ctx, &blogclient.CreatePostRequest{
		UserId:      userId,
		Title:       req.Title,
		Slug:        req.Slug,
		Summary:     req.Summary,
		Content:     req.Content,
		Cover:       req.Cover,
		Status:      req.Status,
		CategoryId:  req.CategoryId,
		PublishedAt: req.PublishedAt,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreatePostResponse{Id: r.Id}, nil
}
