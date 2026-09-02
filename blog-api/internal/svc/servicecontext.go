package svc

import (
	"blog-api/internal/config"
	"blog-rpc/blogclient"
	"user-rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

// ServiceContext blog-api 服务上下文，持有配置和两个 RPC 客户端
type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User // user-rpc gRPC 客户端，处理用户注册/登录/查询
	BlogRpc blogclient.Blog // blog-rpc gRPC 客户端，处理文章增删改查
}

// NewServiceContext 初始化服务上下文，通过 etcd 服务发现创建两个 RPC 客户端
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		BlogRpc: blogclient.NewBlog(zrpc.MustNewClient(c.BlogRpc)),
	}
}
