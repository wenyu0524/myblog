package logic

import (
	"context"
	"time"

	"blog-api/internal/svc"
	"blog-api/internal/types"
	"user-rpc/userclient"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login 用户登录，校验通过后签发 JWT Token
func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 调用 user-rpc 校验用户名和密码
	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	// 构建 JWT claims，将 userId 作为自定义 claim 存入 token
	// 后续请求经 JWT 中间件解析后，userId 会注入到 context 供业务逻辑使用
	now := time.Now().Unix()
	claims := jwt.MapClaims{
		"userId": loginResp.UserId,
		"exp":    now + l.svcCtx.Config.Auth.AccessExpire,
		"iat":    now,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{Token: signedToken}, nil
}
