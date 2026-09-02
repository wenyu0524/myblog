package logic

import (
	"context"

	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPostsLogic {
	return &ListPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListPosts 分页查询文章列表，返回总数和当前页文章数据
func (l *ListPostsLogic) ListPosts(in *blog.ListPostsRequest) (*blog.ListPostsResponse, error) {
	// 查询文章总数，用于分页
	total, err := l.svcCtx.PostModel.Count(l.ctx)
	if err != nil {
		return nil, err
	}

	// 分页查询当前页文章列表
	posts, err := l.svcCtx.PostModel.List(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}

	// 将 model.Post 转换为 protobuf Post，时间戳转为 Unix 秒
	postList := make([]*blog.Post, 0, len(posts))
	for _, p := range posts {
		postList = append(postList, &blog.Post{
			Id:        p.Id,
			UserId:    p.UserId,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: p.CreatedAt.Unix(),
			UpdatedAt: p.UpdatedAt.Unix(),
		})
	}

	return &blog.ListPostsResponse{
		Total: total,
		Posts: postList,
	}, nil
}
