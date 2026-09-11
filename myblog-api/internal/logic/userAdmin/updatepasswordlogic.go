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

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.UpdatePasswordRequest) (resp *types.EmptyResponse, err error) {
	if req == nil {
		return nil, errors.New("请求不能为空")
	}
	userId, err := UserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.UserRpc.UpdatePassword(l.ctx, &userclient.UpdatePasswordRequest{
		UserId:      userId,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return nil, err
	}
	return &types.EmptyResponse{}, nil
}
