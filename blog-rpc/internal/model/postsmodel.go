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
		ListAdmin(ctx context.Context, page, pageSize int64, category, keyword string) (int64, []AdminPost, error)
		GetAdmin(ctx context.Context, id int64) (*AdminPost, error)
		Create(ctx context.Context, data *Posts) (int64, error)
		UpdateAdmin(ctx context.Context, data *Posts) error
		SoftDelete(ctx context.Context, id int64) error
		ListPublic(ctx context.Context, page, pageSize int64, category, keyword string) (int64, []PublicPost, error)
		GetPublic(ctx context.Context, id int64) (*PublicPost, error)
		IncrementView(ctx context.Context, id int64) error
	}
	PublicPost struct {
		Id                                   int64
		Title, Slug, Summary, Content, Cover string
		CategoryId                           int64
		CategoryName                         string
		PublishedAt, CreatedAt, UpdatedAt    int64
	}
	AdminPost struct {
		Id                                           int64
		Title, Slug, Summary, Content, Cover, Status string
		CategoryId                                   int64
		CategoryName                                 string
		ViewCount, PublishedAt, CreatedAt, UpdatedAt int64
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

func (m *customPostsModel) ListAdmin(ctx context.Context, page, pageSize int64, category, keyword string) (int64, []AdminPost, error) {
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
	where, args := "p.deleted_at IS NULL", []any{}
	if category != "" {
		args = append(args, category)
		where += fmt.Sprintf(" AND c.slug=$%d", len(args))
	}
	if keyword != "" {
		args = append(args, "%"+keyword+"%")
		where += fmt.Sprintf(" AND (p.title ILIKE $%d OR p.summary ILIKE $%d)", len(args), len(args))
	}
	var total int64
	if err := m.QueryRowNoCacheCtx(ctx, &total, fmt.Sprintf("SELECT count(*) FROM %s p LEFT JOIN %s c ON c.id=p.category_id WHERE %s", m.table, `"public"."categories"`, where), args...); err != nil {
		return 0, nil, err
	}
	args = append(args, pageSize, off)
	var rows []AdminPost
	q := fmt.Sprintf("SELECT p.id,p.title,p.slug,p.summary,p.content,p.cover,p.status,COALESCE(p.category_id,0),COALESCE(c.name,''),p.view_count,COALESCE(EXTRACT(EPOCH FROM p.published_at),0)::bigint,EXTRACT(EPOCH FROM p.created_at)::bigint,EXTRACT(EPOCH FROM p.updated_at)::bigint FROM %s p LEFT JOIN %s c ON c.id=p.category_id WHERE %s ORDER BY p.created_at DESC,p.id DESC LIMIT $%d OFFSET $%d", m.table, `"public"."categories"`, where, len(args)-1, len(args))
	if err := m.QueryRowsNoCacheCtx(ctx, &rows, q, args...); err != nil {
		return 0, nil, err
	}
	return total, rows, nil
}

func (m *customPostsModel) GetAdmin(ctx context.Context, id int64) (*AdminPost, error) {
	var row AdminPost
	q := fmt.Sprintf("SELECT p.id,p.title,p.slug,p.summary,p.content,p.cover,p.status,COALESCE(p.category_id,0),COALESCE(c.name,''),p.view_count,COALESCE(EXTRACT(EPOCH FROM p.published_at),0)::bigint,EXTRACT(EPOCH FROM p.created_at)::bigint,EXTRACT(EPOCH FROM p.updated_at)::bigint FROM %s p LEFT JOIN %s c ON c.id=p.category_id WHERE p.id=$1 AND p.deleted_at IS NULL", m.table, `"public"."categories"`)
	if err := m.QueryRowNoCacheCtx(ctx, &row, q, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &row, nil
}

func (m *customPostsModel) Create(ctx context.Context, data *Posts) (int64, error) {
	var id int64
	q := fmt.Sprintf("INSERT INTO %s (user_id,category_id,title,slug,summary,content,cover,status,view_count,published_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,0,$9) RETURNING id", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &id, q, data.UserId, data.CategoryId, data.Title, data.Slug, data.Summary, data.Content, data.Cover, data.Status, data.PublishedAt)
	return id, err
}

func (m *customPostsModel) UpdateAdmin(ctx context.Context, data *Posts) error {
	old, err := m.FindOne(ctx, data.Id)
	if err != nil {
		return err
	}
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		q := fmt.Sprintf("UPDATE %s SET title=$2,slug=$3,summary=$4,content=$5,cover=$6,status=$7,category_id=$8,published_at=$9,updated_at=now() WHERE id=$1 AND deleted_at IS NULL", m.table)
		return conn.ExecCtx(ctx, q, data.Id, data.Title, data.Slug, data.Summary, data.Content, data.Cover, data.Status, data.CategoryId, data.PublishedAt)
	}, fmt.Sprintf("%s%v", cachePublicPostsIdPrefix, old.Id), fmt.Sprintf("%s%v", cachePublicPostsSlugPrefix, old.Slug))
	return err
}

func (m *customPostsModel) SoftDelete(ctx context.Context, id int64) error {
	old, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, fmt.Sprintf("UPDATE %s SET deleted_at=now(),updated_at=now() WHERE id=$1 AND deleted_at IS NULL", m.table), id)
	}, fmt.Sprintf("%s%v", cachePublicPostsIdPrefix, old.Id), fmt.Sprintf("%s%v", cachePublicPostsSlugPrefix, old.Slug))
	return err
}
