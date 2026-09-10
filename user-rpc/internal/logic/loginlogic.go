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

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 检查账号密码是否为空
	if in == nil || in.Username == "" || in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "账号和密码不能为空")
	}

	// 根据用户名查用户库
	u, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.Unauthenticated, "用户名或密码错误")
		}
		return nil, status.Error(codes.Internal, "用户查询失败，请稍后重试")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(in.Password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, "用户名或密码错误")
	}

	return &user.LoginResponse{UserId: u.Id, Username: u.Username}, nil
}
