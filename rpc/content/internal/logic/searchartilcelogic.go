package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchArtilceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchArtilceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchArtilceLogic {
	return &SearchArtilceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// * article search *
func (l *SearchArtilceLogic) SearchArtilce(in *content.SearchArtilceRequest) (*content.SearchArtilceResponse, error) {
	//TODO:
	return &content.SearchArtilceResponse{}, nil
}
