package logic

import (
	"context"
	"time"

	"user-rpc/internal/model"
	"user-rpc/internal/svc"
	"user-rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register 用户注册：使用 bcrypt 加密密码后写入数据库
func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 使用 bcrypt 对密码进行哈希加密，不存储明文密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 构建用户记录并插入数据库
	now := time.Now()
	id, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username:  in.Username,
		Password:  string(hashedPassword),
		Email:     in.Email,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return nil, err
	}

	return &user.RegisterResponse{UserId: id}, nil
}
