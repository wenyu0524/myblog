// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package svc

import (
	"blog-rpc/blogclient"
	"myblog-api/internal/config"
	"myblog-api/internal/middleware"
	"user-rpc/userclient"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	BlogRateLimitMiddleware   rest.Middleware
	UploadRateLimitMiddleware rest.Middleware
	UserRateLimitMiddleware   rest.Middleware
	LoginRateLimitMiddleware  rest.Middleware

	UserRpc userclient.User // user-rpc gRPC 客户端，处理用户注册/登录/查询
	BlogRpc blogclient.Blog // blog-rpc gRPC 客户端，处理文章增删改查
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		BlogRateLimitMiddleware:   middleware.NewBlogRateLimitMiddleware(c.RateLimit.BlogPerMinute).Handle,
		UploadRateLimitMiddleware: middleware.NewUploadRateLimitMiddleware(c.RateLimit.UploadPerMinute).Handle,
		UserRateLimitMiddleware:   middleware.NewUserRateLimitMiddleware(c.RateLimit.UserPerMinute).Handle,
		LoginRateLimitMiddleware:  middleware.NewLoginRateLimitMiddleware(c.RateLimit.LoginPerMinute).Handle,

		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		BlogRpc: blogclient.NewBlog(zrpc.MustNewClient(c.BlogRpc)),
	}
}
