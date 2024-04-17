package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArtilceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateArtilceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArtilceLogic {
	return &CreateArtilceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// * article create *
func (l *CreateArtilceLogic) CreateArtilce(in *content.CreateArticleRequest) (*content.CreateArticleResponse, error) {
	// todo: add your logic here and delete this line

	return &content.CreateArticleResponse{}, nil
}
