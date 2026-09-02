package logic

import (
	"context"

	"blog-api/internal/svc"
	"blog-api/internal/types"
	"user-rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMeLogic {
	return &GetMeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetMe 获取当前登录用户信息，从 JWT context 中提取 userId 后调用 user-rpc 查询
func (l *GetMeLogic) GetMe() (resp *types.UserResponse, err error) {
	// 从 context 中提取 JWT 解析后的 userId
	userId, err := getUserId(l.ctx)
	if err != nil {
		return nil, err
	}

	// 调用 user-rpc 获取用户详情
	userResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.GetUserRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserResponse{
		Id:       userResp.UserId,
		Username: userResp.Username,
		Email:    userResp.Email,
	}, nil
}
