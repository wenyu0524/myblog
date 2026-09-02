package logic

import (
	"context"

	"blog-api/internal/svc"
	"blog-api/internal/types"
	"blog-rpc/blogclient"

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

// ListPosts 分页查询文章列表，公开接口无需登录即可访问
func (l *ListPostsLogic) ListPosts(req *types.ListPostsRequest) (resp *types.ListPostsResponse, err error) {
	// 调用 blog-rpc 分页查询文章列表
	listResp, err := l.svcCtx.BlogRpc.ListPosts(l.ctx, &blogclient.ListPostsRequest{
		Page:     int64(req.Page),
		PageSize: int64(req.PageSize),
	})
	if err != nil {
		return nil, err
	}

	// 将 RPC 返回的 Post 列表转换为 API 层的 PostItem 列表
	posts := make([]types.PostItem, 0, len(listResp.Posts))
	for _, p := range listResp.Posts {
		posts = append(posts, types.PostItem{
			Id:        p.Id,
			UserId:    p.UserId,
			Title:     p.Title,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	return &types.ListPostsResponse{
		Total: listResp.Total,
		Posts: posts,
	}, nil
}
