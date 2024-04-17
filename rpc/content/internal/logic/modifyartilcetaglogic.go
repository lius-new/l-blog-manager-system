package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArtilceTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArtilceTagLogic {
	return &ModifyArtilceTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改文章标签
func (l *ModifyArtilceTagLogic) ModifyArtilceTag(in *content.ModifyArticleTagRequest) (*content.ModifyArticleTagResponse, error) {
	// todo: add your logic here and delete this line

	return &content.ModifyArticleTagResponse{}, nil
}
