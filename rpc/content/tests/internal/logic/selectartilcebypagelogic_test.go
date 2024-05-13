package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type SelectArtilceByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectArtilceByPageLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *SelectArtilceByPageLogic {
	return &SelectArtilceByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取文章
func (l *SelectArtilceByPageLogic) SelectArtilceByPage(
	in *content.SelectArticleByPageRequest,
) (*content.SelectArticleByPageResponse, error) {
	if in.PageSize == 0 || in.PageNum == 0 {
		return nil, rpc.ErrRequestParam
	}

	// 根据分页来查询
	articles, total, err := l.svcCtx.ModelWithArticle.FindByPage(
		l.ctx,
		in.PageSize,
		in.PageNum,
		in.HideShow,
	)
	if err != nil {
		return nil, err
	}

	// 封装查询结果为[]*content.SelectArticles类型
	respArticles := make([]*content.SelectArticles, len(articles))

	for _, v := range articles {
		// TODO: 修改返回的tags & 修改返回的COVERS
		respArticles = append(respArticles, &content.SelectArticles{
			Id:     v.ID.Hex(),
			Title:  v.Title,
			Desc:   v.Desc,
			Tags:   v.Tags,
			Covers: v.Covers,
		})
	}

	return &content.SelectArticleByPageResponse{
		Articles: respArticles,
		Total:    total,
	}, nil
}
