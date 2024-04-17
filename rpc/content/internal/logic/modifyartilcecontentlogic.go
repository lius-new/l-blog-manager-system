package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArtilceContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArtilceContentLogic {
	return &ModifyArtilceContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改文章内容
func (l *ModifyArtilceContentLogic) ModifyArtilceContent(in *content.ModifyArticleContentRequest) (*content.ModifyArticleContentResponse, error) {
	// todo: add your logic here and delete this line

	return &content.ModifyArticleContentResponse{}, nil
}
