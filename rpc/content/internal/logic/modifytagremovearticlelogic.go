package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

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
	// todo: add your logic here and delete this line

	return &content.ModifyTagRemoveArticleResponse{}, nil
}
