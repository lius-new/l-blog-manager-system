package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/article"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArtilceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateArtilceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArtilceLogic {
	return &CreateArtilceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// * article create *
func (l *CreateArtilceLogic) CreateArtilce(in *content.CreateArticleRequest) (*content.CreateArticleResponse, error) {
	err := l.svcCtx.ModelWithArticle.Insert(l.ctx, &model.Article{
		Title:    in.Title,
		Desc:     in.Desc,
		Content:  in.Content,
		Covers:   in.Covers,
		Tags:     in.Tags,
		Visiable: in.Visiable,
	})
	if err != nil {
		return nil, err
	}

	return &content.CreateArticleResponse{}, nil
}
