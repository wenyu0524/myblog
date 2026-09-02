package config

import "github.com/zeromicro/go-zero/zrpc"

// Config user-rpc 服务配置，在 zrpc 基础上增加 PostgreSQL 数据源
type Config struct {
	zrpc.RpcServerConf
	DataSource string // PostgreSQL 连接串，指向 user_db
}
