// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package user

import (
	"context"
	"user-rpc/userclient"

	"myblog-api/internal/svc"
	"myblog-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProfileLogic {
	return &GetProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProfileLogic) GetProfile() (resp *types.ProfileResponse, err error) {
	result, err := l.svcCtx.UserRpc.GetProfile(l.ctx, &userclient.Empty{})
	if err != nil {
		return nil, err
	}
	return &types.ProfileResponse{
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
