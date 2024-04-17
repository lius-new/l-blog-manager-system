package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticlesByPageWithBackendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticlesByPageWithBackendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlesByPageWithBackendLogic {
	return &GetArticlesByPageWithBackendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticlesByPageWithBackendLogic) GetArticlesByPageWithBackend(req *types.GetArticlesByPageWithBackendRequest) (resp *types.GetArticlesByPageWithBackendResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
