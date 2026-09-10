// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package blogPublic

import (
	"blog-rpc/blogclient"
	"context"

	"myblog-api/internal/svc"
	"myblog-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPostsLogic {
	return &ListPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPostsLogic) ListPosts(req *types.ListPostsRequest) (resp *types.ListPublicPostsResponse, err error) {
	if req == nil {
		req = &types.ListPostsRequest{}
	}
	r, err := l.svcCtx.BlogRpc.ListPublicPosts(l.ctx, &blogclient.ListPostsRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Category: req.Category,
		Keyword:  req.Keyword},
	)
	if err != nil {
		return nil, err
	}
	resp = &types.ListPublicPostsResponse{Total: r.Total, Posts: make([]types.PublicPostItem, 0, len(r.Posts))}
	for _, p := range r.Posts {
		resp.Posts = append(resp.Posts, types.PublicPostItem{
			Id:           p.Id,
			Title:        p.Title,
			Slug:         p.Slug,
			Summary:      p.Summary,
			Cover:        p.Cover,
			CategoryId:   p.CategoryId,
			CategoryName: p.CategoryName,
			PublishedAt:  p.PublishedAt,
			CreatedAt:    p.CreatedAt,
			UpdatedAt:    p.UpdatedAt,
		})
	}
	return resp, nil
}
