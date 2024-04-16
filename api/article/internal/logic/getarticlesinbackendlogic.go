package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticlesInBackendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticlesInBackendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlesInBackendLogic {
	return &GetArticlesInBackendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticlesInBackendLogic) GetArticlesInBackend(req *types.ViewsReq) (resp *types.ViewsRespInBackend, err error) {
	// todo: add your logic here and delete this line

	return
}
