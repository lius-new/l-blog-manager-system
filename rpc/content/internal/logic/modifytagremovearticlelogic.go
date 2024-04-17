package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/tag"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyTagRemoveArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyTagRemoveArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyTagRemoveArticleLogic {
	return &ModifyTagRemoveArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 从tag中移除article
func (l *ModifyTagRemoveArticleLogic) ModifyTagRemoveArticle(in *content.ModifyTagRemoveArticleRequest) (*content.ModifyTagRemoveArticleResponse, error) {
	currentTag, err := l.svcCtx.ModelWithTag.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	// 判断是否存在
	isExist := -1
	for i := 0; i < len(currentTag.Articles); i++ {
		if currentTag.Articles[i] == in.Article {
			isExist = i
			break
		}
	}
	if isExist == -1 {
		return nil, rpc.ErrNotFound
	}
	articles := append(currentTag.Articles[:isExist], currentTag.Articles[(isExist+1):]...)

	_, err = l.svcCtx.ModelWithTag.Update(l.ctx, &model.Tag{
		ID:       currentTag.ID,
		Articles: articles,
	})
	if err != nil {
		return nil, err
	}

	return &content.ModifyTagRemoveArticleResponse{}, nil
}
