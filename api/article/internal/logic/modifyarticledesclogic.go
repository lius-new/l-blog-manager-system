package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArticleDescLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyArticleDescLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArticleDescLogic {
	return &ModifyArticleDescLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyArticleDescLogic) ModifyArticleDesc(req *types.ModifyArticleDescRequest) (resp *types.ModifyArticleDescResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
