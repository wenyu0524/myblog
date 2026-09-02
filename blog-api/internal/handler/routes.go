package handler

import (
	"net/http"

	"blog-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

// RegisterHandlers 注册所有 HTTP 路由，分为公开路由和 JWT 鉴权路由两组
func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// 公开路由：无需登录即可访问（注册、登录、查看文章）
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/users/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/users/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/posts/get",
				Handler: GetPostHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/posts/list",
				Handler: ListPostsHandler(serverCtx),
			},
		},
	)

	// 鉴权路由：需携带有效 JWT Token，go-zero 自动将 claims 注入 context
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/users/me",
				Handler: GetMeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/posts/create",
				Handler: CreatePostHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/posts/update",
				Handler: UpdatePostHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/posts/delete",
				Handler: DeletePostHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
