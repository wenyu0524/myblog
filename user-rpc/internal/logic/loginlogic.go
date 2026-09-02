package logic

import (
	"context"
	"errors"

	"user-rpc/internal/svc"
	"user-rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/crypto/bcrypt"
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

// Login 用户登录：按用户名查询用户，使用 bcrypt 校验密码
func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 根据用户名查询用户记录
	u, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		// 用户不存在时返回统一的错误信息，避免暴露用户是否存在
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errors.New("invalid username or password")
		}
		return nil, err
	}

	// 使用 bcrypt 对比密码哈希，校验失败返回统一错误信息
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.Password)); err != nil {
		return nil, errors.New("invalid username or password")
	}

	return &user.LoginResponse{
		UserId:   u.Id,
		Username: u.Username,
	}, nil
}
