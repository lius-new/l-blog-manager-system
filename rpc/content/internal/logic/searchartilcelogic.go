package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
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
func (l *SearchArtilceLogic) SearchArtilce(
	in *content.SearchArtilceRequest,
) (*content.SearchArtilceResponse, error) {
	articles, err := l.svcCtx.ModelWithArticle.Search(l.ctx, in.Search)
	if err != nil {
		return nil, err
	}

	resp := make([]*content.SearchArtilce, 0)
	for _, v := range articles {
		resp = append(resp, &content.SearchArtilce{
			Id:    v.ID.Hex(),
			Title: v.Title,
			Desc:  v.Desc,
		})
	}

	return &content.SearchArtilceResponse{
		Articles: resp,
	}, nil
}
