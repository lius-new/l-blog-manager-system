package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArticleTitleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyArticleTitleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArticleTitleLogic {
	return &ModifyArticleTitleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyArticleTitleLogic) ModifyArticleTitle(req *types.ModifyArticleTitleRequest) (resp *types.ModifyArticleTitleResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
