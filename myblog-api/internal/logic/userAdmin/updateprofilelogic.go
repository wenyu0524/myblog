// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package userAdmin

import (
	"context"
	"errors"
	"user-rpc/userclient"

	"myblog-api/internal/svc"
	"myblog-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProfileLogic {
	return &UpdateProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProfileLogic) UpdateProfile(req *types.UpdateProfileRequest) (resp *types.EmptyResponse, err error) {
	if req == nil {
		return nil, errors.New("请求不能为空")
	}
	userId, err := UserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.UserRpc.UpdateProfile(l.ctx, &userclient.UpdateProfileRequest{
		UserId:      userId,
		DisplayName: req.DisplayName,
		Avatar:      req.Avatar,
		Bio:         req.Bio,
		Email:       req.Email,
		Website:     req.Website,
		Github:      req.Github,
	})
	if err != nil {
		return nil, err
	}
	return &types.EmptyResponse{}, nil
}
