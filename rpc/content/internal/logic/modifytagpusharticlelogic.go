package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyTagPushArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyTagPushArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyTagPushArticleLogic {
	return &ModifyTagPushArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加article到tag
func (l *ModifyTagPushArticleLogic) ModifyTagPushArticle(in *content.ModifyTagPushArticleRequest) (*content.ModifyTagPushArticleResponse, error) {
	currentTag, err := l.svcCtx.ModelWithTag.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	articles := append(currentTag.Articles, in.Article)
	_, err = l.svcCtx.ModelWithTag.Update(l.ctx, &model.Tag{
		ID:       currentTag.ID,
		Articles: articles,
	})
	if err != nil {
		return nil, err
	}
	return &content.ModifyTagPushArticleResponse{}, nil
}
