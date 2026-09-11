package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CategoriesModel = (*customCategoriesModel)(nil)

type (
	// CategoriesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCategoriesModel.
	CategoriesModel interface {
		categoriesModel
		ListPublic(ctx context.Context) ([]Categories, error)
		Create(ctx context.Context, data *Categories) (int64, error)
		UpdateAdmin(ctx context.Context, data *Categories) error
	}

	customCategoriesModel struct {
		*defaultCategoriesModel
	}
)

// NewCategoriesModel returns a model for the database table.
func NewCategoriesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CategoriesModel {
	return &customCategoriesModel{
		defaultCategoriesModel: newCategoriesModel(conn, c, opts...),
	}
}

func (m *customCategoriesModel) ListPublic(ctx context.Context) ([]Categories, error) {
	var rows []Categories
	err := m.QueryRowsNoCacheCtx(ctx, &rows, fmt.Sprintf("SELECT id,name,slug,sort,created_at,updated_at FROM %s ORDER BY sort ASC,id ASC", m.table))
	return rows, err
}

func (m *customCategoriesModel) Create(ctx context.Context, data *Categories) (int64, error) {
	var id int64
	err := m.QueryRowNoCacheCtx(ctx, &id, fmt.Sprintf("INSERT INTO %s (name,slug,sort) VALUES ($1,$2,$3) RETURNING id", m.table), data.Name, data.Slug, data.Sort)
	return id, err
}
func (m *customCategoriesModel) UpdateAdmin(ctx context.Context, data *Categories) error {
	old, err := m.FindOne(ctx, data.Id)
	if err != nil {
		return err
	}
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, fmt.Sprintf("UPDATE %s SET name=$2,slug=$3,sort=$4,updated_at=now() WHERE id=$1", m.table), data.Id, data.Name, data.Slug, data.Sort)
	}, fmt.Sprintf("%s%v", cachePublicCategoriesIdPrefix, old.Id), fmt.Sprintf("%s%v", cachePublicCategoriesNamePrefix, old.Name), fmt.Sprintf("%s%v", cachePublicCategoriesSlugPrefix, old.Slug))
	return err
}
