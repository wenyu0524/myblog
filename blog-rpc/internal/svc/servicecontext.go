package svc

import (
	"blog-rpc/internal/config"
	"blog-rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/postgres"
)

// ServiceContext blog-rpc 服务上下文，持有配置和数据库模型
type ServiceContext struct {
	Config     config.Config
	PostsModel model.PostsModel
}

// NewServiceContext 初始化服务上下文，创建 PostgreSQL 连接并构建文章模型
func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.DataSource)
	cacheConf := cache.CacheConf{
		{RedisConf: c.CacheRedis, Weight: 100},
	}
	postsModel := model.NewPostsModel(conn, cacheConf)

	return &ServiceContext{
		Config:     c,
		PostsModel: postsModel,
	}
}
