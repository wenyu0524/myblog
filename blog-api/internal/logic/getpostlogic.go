package logic

import (
	"context"

	"blog-api/internal/svc"
	"blog-api/internal/types"
	"blog-rpc/blogclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLogic {
	return &GetPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetPost 查询单篇文章详情，公开接口无需登录即可访问
func (l *GetPostLogic) GetPost(req *types.GetPostRequest) (resp *types.GetPostResponse, err error) {
	// 调用 blog-rpc 查询文章详情
	getResp, err := l.svcCtx.BlogRpc.GetPost(l.ctx, &blogclient.GetPostRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetPostResponse{
		Id:        getResp.Id,
		UserId:    getResp.UserId,
		Title:     getResp.Title,
		Content:   getResp.Content,
		CreatedAt: getResp.CreatedAt,
		UpdatedAt: getResp.UpdatedAt,
	}, nil
}
