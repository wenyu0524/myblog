package svc

import (
	"user-rpc/internal/config"
	"user-rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.MustNewConn(c.DB)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(conn, c.Cache),
	}
}
