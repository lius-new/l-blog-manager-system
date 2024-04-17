package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

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
	// todo: add your logic here and delete this line

	return &content.ModifyTagPushArticleResponse{}, nil
}
