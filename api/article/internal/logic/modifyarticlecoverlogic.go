package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArticleCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyArticleCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArticleCoverLogic {
	return &ModifyArticleCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyArticleCoverLogic) ModifyArticleCover(req *types.ModifyArticleCoverRequest) (resp *types.ModifyArticleCoverResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
