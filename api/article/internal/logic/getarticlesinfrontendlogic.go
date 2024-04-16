package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticlesInFrontendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticlesInFrontendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlesInFrontendLogic {
	return &GetArticlesInFrontendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticlesInFrontendLogic) GetArticlesInFrontend(req *types.ViewsReq) (resp *types.ViewsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
