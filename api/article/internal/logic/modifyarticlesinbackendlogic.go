package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArticlesInBackendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyArticlesInBackendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArticlesInBackendLogic {
	return &ModifyArticlesInBackendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyArticlesInBackendLogic) ModifyArticlesInBackend(req *types.ReqInBackend) (resp *types.RespInBackend, err error) {
	// todo: add your logic here and delete this line

	return
}
