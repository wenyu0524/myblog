package logic

import (
	"context"
	"errors"
	"net/mail"
	"net/url"
	"unicode/utf8"
	"user-rpc/internal/model"
	"user-rpc/internal/svc"
	"user-rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProfileLogic {
	return &UpdateProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProfileLogic) UpdateProfile(in *user.UpdateProfileRequest) (*user.EmptyResponse, error) {
	if err := validateProfile(in); err != nil {
		return nil, err
	}
	err := l.svcCtx.UserModel.UpdateProfile(l.ctx, &model.Users{
		Id:          in.UserId,
		DisplayName: in.DisplayName,
		Avatar:      in.Avatar,
		Bio:         in.Bio,
		Email:       in.Email,
		Website:     in.Website,
		Github:      in.Github,
	})
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "用户不存在")
		}
		return nil, status.Error(codes.Internal, "用户资料更新失败，请稍后重试")
	}
	return &user.EmptyResponse{}, nil
}

func validateProfile(in *user.UpdateProfileRequest) error {
	if in == nil || in.UserId <= 0 {
		return status.Error(codes.InvalidArgument, "用户ID不能为空")
	}
	if utf8.RuneCountInString(in.DisplayName) > 128 {
		return status.Error(codes.InvalidArgument, "显示名称不能超过128个字符")
	}
	if err := validateHTTPURL(in.Avatar, "头像地址", 2048); err != nil {
		return err
	}
	if utf8.RuneCountInString(in.Bio) > 2000 {
		return status.Error(codes.InvalidArgument, "个人简介不能超过2000个字符")
	}
	if len(in.Email) > 255 {
		return status.Error(codes.InvalidArgument, "邮箱地址不能超过255个字符")
	}
	if in.Email != "" {
		if _, err := mail.ParseAddress(in.Email); err != nil {
			return status.Error(codes.InvalidArgument, "邮箱地址格式不正确")
		}
	}
	if err := validateHTTPURL(in.Website, "个人网站地址", 255); err != nil {
		return err
	}
	return validateHTTPURL(in.Github, "GitHub主页地址", 255)
}

func validateHTTPURL(rawURL, name string, maxLen int) error {
	if rawURL == "" {
		return nil
	}
	if len(rawURL) > maxLen {
		return status.Errorf(codes.InvalidArgument, "%s不能超过%d个字符", name, maxLen)
	}
	u, err := url.Parse(rawURL)
	if err != nil || u.Host == "" || (u.Scheme != "http" && u.Scheme != "https") {
		return status.Errorf(codes.InvalidArgument, "%s格式不正确", name)
	}
	return nil
}
