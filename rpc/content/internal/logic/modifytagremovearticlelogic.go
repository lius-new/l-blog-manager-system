package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/tag"
)

type ModifyTagRemoveArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyTagRemoveArticleLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *ModifyTagRemoveArticleLogic {
	return &ModifyTagRemoveArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 从tag中移除article
func (l *ModifyTagRemoveArticleLogic) ModifyTagRemoveArticle(
	in *content.ModifyTagRemoveArticleRequest,
) (*content.ModifyTagRemoveArticleResponse, error) {

	// 判断是否存在并查询出结果
	currentTag, err := l.svcCtx.ModelWithTag.FindOne(l.ctx, in.Id)
	if err == rpc.ErrNotFound || currentTag == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// 遍历当前tag上所有的articleid, 如果存在指定articleId那么就设置isExist 为指定下标
	isExist := -1
	for i := 0; i < len(currentTag.Articles); i++ {
		if currentTag.Articles[i] == in.Article {
			isExist = i
			break
		}
	}
	// 未在tag上查询到指定article
	if isExist == -1 {
		return nil, rpc.ErrNotFound
	}

	// 在tag的articles中移除指定articleId
	articles := append(currentTag.Articles[:isExist], currentTag.Articles[(isExist+1):]...)

	// 更新数据
	_, err = l.svcCtx.ModelWithTag.Update(l.ctx, &model.Tag{
		ID:       currentTag.ID,
		Articles: articles,
	})
	if err != nil {
		return nil, err
	}

	return &content.ModifyTagRemoveArticleResponse{}, nil
}
