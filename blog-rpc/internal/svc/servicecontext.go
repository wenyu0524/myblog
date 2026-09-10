package svc

import (
	"blog-rpc/internal/config"
	"blog-rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type ServiceContext struct {
	Config          config.Config
	CategoriesModel model.CategoriesModel
	PostsModel      model.PostsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.MustNewConn(c.DB)

	return &ServiceContext{
		Config:          c,
		CategoriesModel: model.NewCategoriesModel(conn, c.Cache),
		PostsModel:      model.NewPostsModel(conn, c.Cache),
	}
}
