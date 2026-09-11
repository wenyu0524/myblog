// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package blogAdmin

import (
	"blog-rpc/blogclient"
	"context"
	"errors"
	userAdmin "myblog-api/internal/logic/userAdmin"

	"myblog-api/internal/svc"
	"myblog-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImageLogic {
	return &UploadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImageLogic) UploadImage(req *types.UploadImageRequest) (resp *types.UploadImageResponse, err error) {
	if req == nil {
		return nil, errors.New("请求不能为空")
	}
	id, err := userAdmin.UserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	r, err := l.svcCtx.BlogRpc.UploadImage(l.ctx, &blogclient.UploadImageRequest{UserId: id, Filename: req.Filename, ContentType: req.ContentType, Content: req.Content})
	if err != nil {
		return nil, err
	}
	return &types.UploadImageResponse{Url: r.Url}, nil
}
