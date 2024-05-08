package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArticleTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyArticleTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArticleTagLogic {
	return &ModifyArticleTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyArticleTagLogic) ModifyArticleTag(req *types.ModifyArticleTagRequest) (resp *types.ModifyArticleTagResponse, err error) {
	// todo: add your logic here and delete this line

	return
}