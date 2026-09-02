package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// Config blog-api HTTP 服务配置
type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string // JWT 签名密钥
		AccessExpire int64  // JWT 过期时间（秒）
	}
	UserRpc zrpc.RpcClientConf // user-rpc gRPC 客户端配置（通过 etcd 服务发现）
	BlogRpc zrpc.RpcClientConf // blog-rpc gRPC 客户端配置（通过 etcd 服务发现）
}
