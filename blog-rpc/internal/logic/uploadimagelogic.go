package logic

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UploadImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImageLogic {
	return &UploadImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadImageLogic) UploadImage(in *blog.UploadImageRequest) (*blog.UploadImageResponse, error) {
	if in == nil || in.UserId <= 0 || len(in.Content) == 0 {
		return nil, status.Error(codes.InvalidArgument, "用户ID和图片内容不能为空")
	}
	if int64(len(in.Content)) > l.svcCtx.Config.Upload.MaxBytes {
		return nil, status.Error(codes.InvalidArgument, "图片大小超过限制")
	}
	if len(in.Filename) > 100 {
		return nil, status.Error(codes.InvalidArgument, "文件名过长")
	}
	kind := http.DetectContentType(in.Content)
	ext := map[string]string{"image/jpeg": ".jpg", "image/png": ".png", "image/gif": ".gif", "image/webp": ".webp"}[kind]
	if ext == "" || (!strings.HasPrefix(in.ContentType, "image/") && in.ContentType != "") {
		return nil, status.Error(codes.InvalidArgument, "仅支持 JPEG、PNG、GIF 或 WebP 图片")
	}
	var b [12]byte
	if _, err := rand.Read(b[:]); err != nil {
		return nil, status.Error(codes.Internal, "图片保存失败")
	}
	name := hex.EncodeToString(b[:]) + ext
	if err := os.MkdirAll(l.svcCtx.Config.Upload.Dir, 0755); err != nil {
		return nil, status.Error(codes.Internal, "图片保存失败")
	}
	if err := os.WriteFile(filepath.Join(l.svcCtx.Config.Upload.Dir, name), in.Content, 0644); err != nil {
		return nil, status.Error(codes.Internal, "图片保存失败")
	}
	return &blog.UploadImageResponse{Url: strings.TrimRight(l.svcCtx.Config.Upload.BaseUrl, "/") + "/" + name}, nil
}
