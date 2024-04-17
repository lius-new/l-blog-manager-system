package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByTagIdWithViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByTagIdWithViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByTagIdWithViewLogic {
	return &GetArticleByTagIdWithViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByTagIdWithViewLogic) GetArticleByTagIdWithView(req *types.GetArticleByTagIdWithViewRequest) (resp *types.GetArticleByTagIdWithViewResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
