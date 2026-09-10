package model

import (
	"context"
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
