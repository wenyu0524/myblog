// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package main

import (
	"flag"
	"fmt"
	"net/http"

	"myblog-api/internal/config"
	"myblog-api/internal/handler"
	"myblog-api/internal/middleware"
	"myblog-api/internal/svc"
	"myblog-api/response"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/myblog-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	httpx.SetErrorHandler(func(err error) (int, any) {
		return http.StatusBadRequest, response.Body{Code: -1, Message: err.Error(), Data: nil}
	})

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		logx.WithContext(r.Context()).WithFields(logx.Field("method", r.Method), logx.Field("path", r.URL.Path), logx.Field("client_ip", r.RemoteAddr)).Errorf("unauthorized request: %v", err)
		httpx.WriteJson(w, http.StatusUnauthorized, response.Body{Code: -1, Message: err.Error()})
	}))
	server.Use(middleware.RequestLog)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
