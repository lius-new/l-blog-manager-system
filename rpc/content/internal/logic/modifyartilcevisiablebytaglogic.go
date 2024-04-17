package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArtilceVisiableByTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceVisiableByTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArtilceVisiableByTagLogic {
	return &ModifyArtilceVisiableByTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据tag修改文章的可见性
func (l *ModifyArtilceVisiableByTagLogic) ModifyArtilceVisiableByTag(in *content.ModifyArticleVisiableByTagRequest) (*content.ModifyArticleVisiableByTagResponse, error) {
	// todo: add your logic here and delete this line

	return &content.ModifyArticleVisiableByTagResponse{}, nil
}
