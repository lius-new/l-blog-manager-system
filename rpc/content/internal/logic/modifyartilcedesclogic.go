package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArtilceDescLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceDescLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArtilceDescLogic {
	return &ModifyArtilceDescLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改文章描述
func (l *ModifyArtilceDescLogic) ModifyArtilceDesc(in *content.ModifyArticleDescRequest) (*content.ModifyArticleDescResponse, error) {
	// todo: add your logic here and delete this line

	return &content.ModifyArticleDescResponse{}, nil
}
