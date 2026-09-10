// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package blogPublic

import (
	"blog-rpc/blogclient"
	"context"
	"fmt"

	"myblog-api/internal/svc"
	"myblog-api/internal/types"

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

func (l *GetPostLogic) GetPost(req *types.GetPostRequest) (resp *types.PublicPostResponse, err error) {
	if req == nil {
		return nil, fmt.Errorf("请求不能为空")
	}
	r, err := l.svcCtx.BlogRpc.GetPublicPost(l.ctx, &blogclient.GetPostRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.PublicPostResponse{
		Id:           r.Id,
		Title:        r.Title,
		Slug:         r.Slug,
		Summary:      r.Summary,
		Content:      r.Content,
		Cover:        r.Cover,
		CategoryId:   r.CategoryId,
		CategoryName: r.CategoryName,
		PublishedAt:  r.PublishedAt,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}, nil
}
