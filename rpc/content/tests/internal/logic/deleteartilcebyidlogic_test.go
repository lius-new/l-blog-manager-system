package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type DeleteArtilceByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteArtilceByIdLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *DeleteArtilceByIdLogic {
	return &DeleteArtilceByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据删除文章
func (l *DeleteArtilceByIdLogic) DeleteArtilceById(
	in *content.DeleteArticleRequest,
) (*content.DeleteArticleResponse, error) {
	count, err := l.svcCtx.ModelWithArticle.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &content.DeleteArticleResponse{Count: count}, nil
}
