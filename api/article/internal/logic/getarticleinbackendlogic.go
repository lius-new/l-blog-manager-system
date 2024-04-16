package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleInBackendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleInBackendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleInBackendLogic {
	return &GetArticleInBackendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleInBackendLogic) GetArticleInBackend() (resp *types.ArticleBackend, err error) {
	// todo: add your logic here and delete this line

	return
}
