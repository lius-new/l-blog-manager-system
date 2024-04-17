package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByIdWithViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByIdWithViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByIdWithViewLogic {
	return &GetArticleByIdWithViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByIdWithViewLogic) GetArticleByIdWithView(req *types.GetArticleByIdWithViewRequest) (resp *types.GetArticleByIdWithViewResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
