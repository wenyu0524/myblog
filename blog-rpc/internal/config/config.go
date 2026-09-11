package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB     sqlx.SqlConf
	Cache  cache.CacheConf
	Upload struct {
		Dir      string
		BaseUrl  string
		MaxBytes int64
	}
}
