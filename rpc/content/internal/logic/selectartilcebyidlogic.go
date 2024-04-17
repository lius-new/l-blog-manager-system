package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectArtilceByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectArtilceByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectArtilceByIdLogic {
	return &SelectArtilceByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// * article select *
func (l *SelectArtilceByIdLogic) SelectArtilceById(in *content.SelectArticleByIdRequest) (*content.SelectArticle, error) {
	// todo: add your logic here and delete this line

	return &content.SelectArticle{}, nil
}
