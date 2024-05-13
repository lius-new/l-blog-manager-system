package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type ExistArtilceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExistArtilceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExistArtilceLogic {
	return &ExistArtilceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// * article exist *
func (l *ExistArtilceLogic) ExistArtilce(
	in *content.ExistArtilceRequest,
) (*content.ExistArtilceResponse, error) {

	findArticle, err := l.svcCtx.ModelWithArticle.FindOne(l.ctx, in.Id)

	if err == rpc.ErrNotFound || findArticle == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &content.ExistArtilceResponse{
		Exist: true,
	}, nil
}
