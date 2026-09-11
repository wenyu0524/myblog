// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// User RPC 配置
	UserRpc zrpc.RpcClientConf

	// Blog RPC 配置
	BlogRpc zrpc.RpcClientConf

	// 限流配置
	RateLimit struct {
		UserPerMinute   int
		LoginPerMinute  int
		BlogPerMinute   int
		UploadPerMinute int
	}

	// JWT 配置
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
