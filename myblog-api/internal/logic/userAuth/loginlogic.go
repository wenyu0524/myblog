// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package userAuth

import (
	"context"
	"time"
	"user-rpc/userclient"

	"myblog-api/internal/svc"
	"myblog-api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 调用user-rpc校验用户名和密码
	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	expiresAt := now + l.svcCtx.Config.Auth.AccessExpire
	claims := jwt.MapClaims{
		"userId":   loginResp.UserId,
		"username": loginResp.Username,
		"iat":      now,
		"exp":      expiresAt,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{Token: signedToken, ExpiresAt: expiresAt}, nil
}
