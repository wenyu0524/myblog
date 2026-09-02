package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const postCacheKeyPrefix = "blog:post:"

// Post 文章表数据模型
type Post struct {
	Id        int64     `db:"id"`
	UserId    int64     `db:"user_id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// PostModel 文章模型接口，定义文章的数据库 CRUD 操作
type PostModel interface {
	// Insert 新增文章，返回文章 ID
	Insert(ctx context.Context, data *Post) (int64, error)
	// FindOne 根据文章 ID 查询文章详情
	FindOne(ctx context.Context, id int64) (*Post, error)
	// List 分页查询文章列表（按 ID 倒序）
	List(ctx context.Context, page, pageSize int64) ([]*Post, error)
	// Count 查询文章总数（用于分页）
	Count(ctx context.Context) (int64, error)
	// Update 更新文章，返回受影响行数（0 表示无权操作或文章不存在）
	Update(ctx context.Context, data *Post) (int64, error)
	// Delete 删除文章，同时校验 user_id 确保仅作者可删除
	Delete(ctx context.Context, id, userId int64) (int64, error)
}

type postModel struct {
	sqlc.CachedConn
	conn sqlx.SqlConn
}

// NewPostModel 创建文章模型实例
func NewPostModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PostModel {
	return &postModel{CachedConn: sqlc.NewConn(conn, c, opts...), conn: conn}
}

// Insert 新增文章，使用 RETURNING id 获取自增主键
func (m *postModel) Insert(ctx context.Context, data *Post) (int64, error) {
	query := `INSERT INTO posts (user_id, title, content, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int64
	err := m.conn.QueryRowCtx(ctx, &id, query,
		data.UserId, data.Title, data.Content, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// FindOne 根据文章 ID 查询单篇文章详情
func (m *postModel) FindOne(ctx context.Context, id int64) (*Post, error) {
	query := `SELECT id, user_id, title, content, created_at, updated_at
		FROM posts WHERE id = $1`
	var post Post
	err := m.QueryRowCtx(ctx, &post, postCacheKeyPrefix+fmt.Sprint(id), func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, sqlx.ErrNotFound
		}
		return nil, err
	}
	return &post, nil
}

// List 分页查询文章列表，按 ID 倒序排列，offset 从 0 开始
func (m *postModel) List(ctx context.Context, page, pageSize int64) ([]*Post, error) {
	offset := (page - 1) * pageSize
	query := `SELECT id, user_id, title, content, created_at, updated_at
		FROM posts ORDER BY id DESC LIMIT $1 OFFSET $2`
	var posts []*Post
	err := m.conn.QueryRowsCtx(ctx, &posts, query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// Count 查询文章总记录数，用于分页响应中的 total 字段
func (m *postModel) Count(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM posts`
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Update 更新文章，WHERE 条件同时校验 id 和 user_id，确保仅作者可修改
func (m *postModel) Update(ctx context.Context, data *Post) (int64, error) {
	query := `UPDATE posts SET title = $1, content = $2, updated_at = $3
		WHERE id = $4 AND user_id = $5`
	res, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, data.Title, data.Content, data.UpdatedAt, data.Id, data.UserId)
	}, postCacheKeyPrefix+fmt.Sprint(data.Id))
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// Delete 删除文章，WHERE 条件同时校验 id 和 user_id，确保仅作者可删除
func (m *postModel) Delete(ctx context.Context, id, userId int64) (int64, error) {
	query := `DELETE FROM posts WHERE id = $1 AND user_id = $2`
	res, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, id, userId)
	}, postCacheKeyPrefix+fmt.Sprint(id))
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
