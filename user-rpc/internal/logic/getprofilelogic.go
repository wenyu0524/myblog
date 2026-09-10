package logic

import (
	"context"

	"user-rpc/internal/svc"
	"user-rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProfileLogic {
	return &GetProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProfileLogic) GetProfile(in *user.Empty) (*user.ProfileResponse, error) {
	// 按作者ID查询用户信息
	profile, err := l.svcCtx.UserModel.FindOne(l.ctx, l.svcCtx.Config.OwnerUserID)
	if err != nil {
		return nil, err
	}

	return &user.ProfileResponse{
		Id:          profile.Id,
		Username:    profile.Username,
		DisplayName: profile.DisplayName,
		Avatar:      profile.Avatar,
		Bio:         profile.Bio,
		Email:       profile.Email,
		Website:     profile.Website,
		Github:      profile.Github,
	}, nil
}
