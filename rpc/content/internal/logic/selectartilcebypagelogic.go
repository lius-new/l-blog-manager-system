package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectArtilceByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectArtilceByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectArtilceByPageLogic {
	return &SelectArtilceByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取文章
func (l *SelectArtilceByPageLogic) SelectArtilceByPage(in *content.SelectArticleByPageRequest) (*content.SelectArticleByPageResponse, error) {
	// TODO:

	return &content.SelectArticleByPageResponse{}, nil
}
