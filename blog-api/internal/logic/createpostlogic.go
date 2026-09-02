package logic

import (
	"context"

	"blog-api/internal/svc"
	"blog-api/internal/types"
	"blog-rpc/blogclient"

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

// CreatePost 创建文章，从 context 提取 userId 作为作者，转发给 blog-rpc 写入数据库
func (l *CreatePostLogic) CreatePost(req *types.CreatePostRequest) (resp *types.CreatePostResponse, err error) {
	// 从 JWT context 中获取当前登录用户 ID，作为文章作者
	userId, err := getUserId(l.ctx)
	if err != nil {
		return nil, err
	}

	// 调用 blog-rpc 创建文章
	createResp, err := l.svcCtx.BlogRpc.CreatePost(l.ctx, &blogclient.CreatePostRequest{
		UserId:  userId,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreatePostResponse{Id: createResp.Id}, nil
}
