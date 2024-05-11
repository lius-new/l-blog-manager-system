package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type SelectArtilceByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectArtilceByIdLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *SelectArtilceByIdLogic {
	return &SelectArtilceByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// * article select *
func (l *SelectArtilceByIdLogic) SelectArtilceById(
	in *content.SelectArticleByIdRequest,
) (*content.SelectArticle, error) {
	// 查询指定id的文章
	findArticle, err := l.svcCtx.ModelWithArticle.FindOne(l.ctx, in.Id)
	if err == rpc.ErrNotFound || findArticle == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}
  // TODO: 修改返回的tags & 修改返回的COVERS

	return &content.SelectArticle{
		Id:      findArticle.ID.Hex(),
		Title:   findArticle.Title,
		Desc:    findArticle.Desc,
		Content: findArticle.Content,
		Tags:    findArticle.Tags,
		Covers:  findArticle.Covers,
	}, nil
}
