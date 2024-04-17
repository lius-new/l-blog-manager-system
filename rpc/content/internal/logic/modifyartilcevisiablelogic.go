package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArtilceVisiableLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceVisiableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArtilceVisiableLogic {
	return &ModifyArtilceVisiableLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改文章的可见性
func (l *ModifyArtilceVisiableLogic) ModifyArtilceVisiable(in *content.ModifyArticleVisiableRequest) (*content.ModifyArticleVisiableResponse, error) {
	// todo: add your logic here and delete this line

	return &content.ModifyArticleVisiableResponse{}, nil
}
