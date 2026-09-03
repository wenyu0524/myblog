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
	// 使用 c.DataSource 中的数据库连接配置，创建 PostgreSQL 数据库连接对象。
	conn := postgres.New(c.DataSource)
	// 创建 Redis 缓存配置。
	// CacheConf 是一个缓存配置列表，这里配置了一个 Redis 缓存节点。
	cacheConf := cache.CacheConf{
		// 使用 c.CacheRedis 中的 Redis 配置，并设置权重为 100。
		// Weight 通常用于多个缓存节点之间的权重分配，这里只有一个节点，所以设置为 100。
		{RedisConf: c.CacheRedis, Weight: 100},
	}
	// 创建 Posts 数据模型。
	// conn：使用上面创建的 PostgreSQL 数据库连接。
	// cacheConf：使用上面配置的 Redis 缓存。
	// postsModel：创建完成后，用于操作 posts 相关的数据。
	postsModel := model.NewPostsModel(conn, cacheConf)

	return &ServiceContext{
		Config:     c,
		PostsModel: postsModel,
	}
}
