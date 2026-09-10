// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package userAdmin

import (
	"context"
	"user-rpc/userclient"

	"myblog-api/internal/svc"
	"myblog-api/internal/types"

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

func (l *GetMeLogic) GetMe() (resp *types.UserResponse, err error) {
	userId, err := userIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	result, err := l.svcCtx.UserRpc.GetMe(l.ctx, &userclient.GetMeRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return &types.UserResponse{
		Id:          result.Id,
		Username:    result.Username,
		DisplayName: result.DisplayName,
		Avatar:      result.Avatar,
		Bio:         result.Bio,
		Email:       result.Email,
		Website:     result.Website,
		Github:      result.Github,
	}, nil
}
