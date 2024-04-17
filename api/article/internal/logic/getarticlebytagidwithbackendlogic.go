package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByTagIdWithBackendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByTagIdWithBackendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByTagIdWithBackendLogic {
	return &GetArticleByTagIdWithBackendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByTagIdWithBackendLogic) GetArticleByTagIdWithBackend(req *types.GetArticleByTagIdWithBackendRequest) (resp *types.GetArticleByTagIdWithBackendResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
