package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostsModel = (*customPostsModel)(nil)

type (
	// PostsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostsModel.
	PostsModel interface {
		postsModel
		ListPublic(ctx context.Context, page, pageSize int64, category, keyword string) (int64, []PublicPost, error)
		GetPublic(ctx context.Context, id int64) (*PublicPost, error)
		IncrementView(ctx context.Context, id int64) error
	}
	PublicPost struct {
		Id, CategoryId, PublishedAt, CreatedAt, UpdatedAt  int64
		Title, Slug, Summary, Content, Cover, CategoryName string
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

func (m *customPostsModel) ListPublic(ctx context.Context, page, pageSize int64, category, keyword string) (int64, []PublicPost, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	off := (page - 1) * pageSize
	var total int64
	var rows []PublicPost
	where := "p.status = 'published' AND p.deleted_at IS NULL"
	args := []any{}
	if category != "" {
		args = append(args, category)
		where += fmt.Sprintf(" AND c.slug = $%d", len(args))
	}
	if keyword != "" {
		args = append(args, "%"+keyword+"%")
		where += fmt.Sprintf(" AND (p.title ILIKE $%d OR p.summary ILIKE $%d)", len(args), len(args))
	}
	countQ := fmt.Sprintf("SELECT count(*) FROM %s p LEFT JOIN %s c ON c.id=p.category_id WHERE %s", m.table, `"public"."categories"`, where)
	if err := m.QueryRowNoCacheCtx(ctx, &total, countQ, args...); err != nil {
		return 0, nil, err
	}
	args = append(args, pageSize, off)
	q := fmt.Sprintf("SELECT p.id,p.title,p.slug,p.summary,p.content,p.cover,COALESCE(p.category_id,0),COALESCE(c.name,''),COALESCE(EXTRACT(EPOCH FROM p.published_at),0)::bigint,EXTRACT(EPOCH FROM p.created_at)::bigint,EXTRACT(EPOCH FROM p.updated_at)::bigint FROM %s p LEFT JOIN %s c ON c.id=p.category_id WHERE %s ORDER BY p.published_at DESC NULLS LAST,p.id DESC LIMIT $%d OFFSET $%d", m.table, `"public"."categories"`, where, len(args)-1, len(args))
	if err := m.QueryRowsNoCacheCtx(ctx, &rows, q, args...); err != nil {
		return 0, nil, err
	}
	return total, rows, nil
}

func (m *customPostsModel) GetPublic(ctx context.Context, id int64) (*PublicPost, error) {
	var row PublicPost
	q := fmt.Sprintf("SELECT p.id,p.title,p.slug,p.summary,p.content,p.cover,COALESCE(p.category_id,0),COALESCE(c.name,''),COALESCE(EXTRACT(EPOCH FROM p.published_at),0)::bigint,EXTRACT(EPOCH FROM p.created_at)::bigint,EXTRACT(EPOCH FROM p.updated_at)::bigint FROM %s p LEFT JOIN %s c ON c.id=p.category_id WHERE p.id=$1 AND p.status='published' AND p.deleted_at IS NULL", m.table, `"public"."categories"`)
	if err := m.QueryRowNoCacheCtx(ctx, &row, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &row, nil
}

func (m *customPostsModel) IncrementView(ctx context.Context, id int64) error {
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, fmt.Sprintf("UPDATE %s SET view_count=view_count+1 WHERE id=$1 AND status='published' AND deleted_at IS NULL", m.table), id)
	}, fmt.Sprintf("%s%v", cachePublicPostsIdPrefix, id))
	return err
}
