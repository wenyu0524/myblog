package logic

import (
	"context"
	"errors"
	"user-rpc/internal/model"
	"user-rpc/internal/svc"
	"user-rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdatePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(in *user.UpdatePasswordRequest) (*user.EmptyResponse, error) {
	if in == nil || in.UserId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "用户ID不能为空")
	}
	if in.OldPassword == "" || in.NewPassword == "" {
		return nil, status.Error(codes.InvalidArgument, "旧密码和新密码不能为空")
	}
	if len(in.NewPassword) < 12 {
		return nil, status.Error(codes.InvalidArgument, "新密码不能少于12个字符")
	}
	if len(in.NewPassword) > 72 {
		return nil, status.Error(codes.InvalidArgument, "新密码不能超过72个字符")
	}
	if in.OldPassword == in.NewPassword {
		return nil, status.Error(codes.InvalidArgument, "新密码不能与旧密码相同")
	}

	u, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "用户不存在")
		}
		return nil, status.Error(codes.Internal, "用户查询失败，请稍后重试")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(in.OldPassword)); err != nil {
		return nil, status.Error(codes.Unauthenticated, "旧密码错误")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "密码更新失败，请稍后重试")
	}
	if err := l.svcCtx.UserModel.UpdatePassword(l.ctx, in.UserId, string(hash)); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "用户不存在")
		}
		return nil, status.Error(codes.Internal, "密码更新失败，请稍后重试")
	}
	return &user.EmptyResponse{}, nil
}
