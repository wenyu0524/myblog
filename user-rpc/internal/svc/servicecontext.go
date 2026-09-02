package svc

import (
	"user-rpc/internal/config"
	"user-rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/postgres"
)

// ServiceContext user-rpc 服务上下文，持有配置和数据库模型
type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

// NewServiceContext 初始化服务上下文，创建 PostgreSQL 连接并构建用户模型
func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn),
	}
}
