package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArtilceTitleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceTitleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArtilceTitleLogic {
	return &ModifyArtilceTitleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// * article modify *
func (l *ModifyArtilceTitleLogic) ModifyArtilceTitle(in *content.ModifyArticleTitleRequest) (*content.ModifyArticleTitleResponse, error) {
	// todo: add your logic here and delete this line

	return &content.ModifyArticleTitleResponse{}, nil
}
