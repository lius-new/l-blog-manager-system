package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectArtilceByTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectArtilceByTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectArtilceByTagLogic {
	return &SelectArtilceByTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据tag获取文章
func (l *SelectArtilceByTagLogic) SelectArtilceByTag(in *content.SelectArticleByTagRequest) (*content.SelectArticleByTagResponse, error) {
	// todo: add your logic here and delete this line

	return &content.SelectArticleByTagResponse{}, nil
}
