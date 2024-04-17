package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticlesByPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticlesByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlesByPageLogic {
	return &GetArticlesByPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticlesByPageLogic) GetArticlesByPage(req *types.GetArticleByPageRequest) (resp *types.GetArticleByPageWithViewResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
