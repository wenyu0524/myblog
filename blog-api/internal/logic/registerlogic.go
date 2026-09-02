package logic

import (
	"context"

	"blog-api/internal/svc"
	"blog-api/internal/types"
	"user-rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register 用户注册，将请求转发给 user-rpc 进行用户创建
func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// 调用 user-rpc 注册接口，密码加密在 RPC 层完成
	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &userclient.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResponse{UserId: registerResp.UserId}, nil
}
