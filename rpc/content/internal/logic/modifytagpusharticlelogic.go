package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/tag"
)

type ModifyTagPushArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyTagPushArticleLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *ModifyTagPushArticleLogic {
	return &ModifyTagPushArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加article到tag
func (l *ModifyTagPushArticleLogic) ModifyTagPushArticle(
	in *content.ModifyTagPushArticleRequest,
) (*content.ModifyTagPushArticleResponse, error) {
	// 找到指定tag
	currentTag, err := l.svcCtx.ModelWithTag.FindOne(l.ctx, in.Id)

	// 判断是否存在错误或者tag是否存在
	if err == rpc.ErrNotFound || currentTag == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// 将articleId添加到currentTag.Articles中并更新数据
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
