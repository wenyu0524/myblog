package svc

import (
	"blog-api/internal/config"
	"blog-rpc/blogclient"
	"user-rpc/userclient"

	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

// ServiceContext blog-api 服务上下文，持有配置和两个 RPC 客户端
type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User // user-rpc gRPC 客户端，处理用户注册/登录/查询
	BlogRpc blogclient.Blog // blog-rpc gRPC 客户端，处理文章增删改查

	RegisterLimiter *limit.PeriodLimit
	LoginLimiter    *limit.PeriodLimit
	GetPostLimiter  *limit.PeriodLimit
	ListLimiter     *limit.PeriodLimit
	WriteLimiter    *limit.PeriodLimit
}

// NewServiceContext 初始化服务上下文，通过 etcd 服务发现创建两个 RPC 客户端
func NewServiceContext(c config.Config) *ServiceContext {
	redisClient := redis.MustNewRedis(c.RateLimitRedis)

	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		BlogRpc: blogclient.NewBlog(zrpc.MustNewClient(c.BlogRpc)),

		RegisterLimiter: limit.NewPeriodLimit(60, c.RateLimit.RegisterPerMinute, redisClient, "rate:register:"),
		LoginLimiter:    limit.NewPeriodLimit(60, c.RateLimit.LoginPerMinute, redisClient, "rate:login:"),
		GetPostLimiter:  limit.NewPeriodLimit(60, c.RateLimit.GetPostPerMinute, redisClient, "rate:get-post:"),
		ListLimiter:     limit.NewPeriodLimit(60, c.RateLimit.ListPostsPerMinute, redisClient, "rate:list-posts:"),
		WriteLimiter:    limit.NewPeriodLimit(60, c.RateLimit.WritePostPerMinute, redisClient, "rate:write-post:"),
	}
}
