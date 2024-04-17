package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArticleVisiableLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyArticleVisiableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArticleVisiableLogic {
	return &ModifyArticleVisiableLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyArticleVisiableLogic) ModifyArticleVisiable(req *types.ModifyArticleVisiableRequest) (resp *types.ModifyArticleVisiableResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
