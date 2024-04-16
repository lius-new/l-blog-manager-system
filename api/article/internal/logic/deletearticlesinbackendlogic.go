package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticlesInBackendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticlesInBackendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticlesInBackendLogic {
	return &DeleteArticlesInBackendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticlesInBackendLogic) DeleteArticlesInBackend() (resp *types.RespInBackend, err error) {
	// todo: add your logic here and delete this line

	return
}
