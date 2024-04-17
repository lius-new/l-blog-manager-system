package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByIdWithBackendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByIdWithBackendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByIdWithBackendLogic {
	return &GetArticleByIdWithBackendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByIdWithBackendLogic) GetArticleByIdWithBackend(req *types.GetArticleByIdWithBackendRequest) (resp *types.GetArticleByIdWithBackendResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
