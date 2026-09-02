package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostsModel = (*customPostsModel)(nil)

type (
	// PostsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostsModel.
	PostsModel interface {
		postsModel
		InsertReturningID(ctx context.Context, data *Posts) (int64, error)
		List(ctx context.Context, page, pageSize int64) ([]*Posts, error)
		Count(ctx context.Context) (int64, error)
	}

	customPostsModel struct {
		*defaultPostsModel
	}
)

// NewPostsModel returns a model for the database table.
func NewPostsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PostsModel {
	return &customPostsModel{
		defaultPostsModel: newPostsModel(conn, c, opts...),
	}
}

// InsertReturningID inserts a post and returns its PostgreSQL-generated ID.
func (m *customPostsModel) InsertReturningID(ctx context.Context, data *Posts) (int64, error) {
	query := "insert into " + m.table + " (user_id, title, content) values ($1, $2, $3) returning id"
	var id int64
	err := m.QueryRowNoCacheCtx(ctx, &id, query, data.UserId, data.Title, data.Content)
	return id, err
}

// List returns a page of posts ordered by newest ID first.
func (m *customPostsModel) List(ctx context.Context, page, pageSize int64) ([]*Posts, error) {
	offset := (page - 1) * pageSize
	query := "select " + postsRows + " from " + m.table + " order by id desc limit $1 offset $2"
	var posts []*Posts
	err := m.QueryRowsNoCacheCtx(ctx, &posts, query, pageSize, offset)
	return posts, err
}

// Count returns the total number of posts.
func (m *customPostsModel) Count(ctx context.Context) (int64, error) {
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, "select count(*) from "+m.table)
	return count, err
}
