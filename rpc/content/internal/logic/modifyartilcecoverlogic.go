package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArtilceCoverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArtilceCoverLogic {
	return &ModifyArtilceCoverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改文章Cover
func (l *ModifyArtilceCoverLogic) ModifyArtilceCover(in *content.ModifyArticleCoverRequest) (*content.ModifyArticleCoverResponse, error) {
	// todo: add your logic here and delete this line

	return &content.ModifyArticleCoverResponse{}, nil
}
