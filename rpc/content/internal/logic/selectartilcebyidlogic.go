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
	article, err := l.svcCtx.ModelWithArticle.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &content.SelectArticle{
		Id:      article.ID.Hex(),
		Title:   article.Title,
		Desc:    article.Desc,
		Content: article.Content,
		Tags:    article.Tags,
		Covers:  article.Covers,
	}, nil
}
