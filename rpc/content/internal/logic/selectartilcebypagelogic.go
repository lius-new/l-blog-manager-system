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
		in.PageNum,
		in.PageSize,
		in.HideShow,
	)
	if err != nil {
		return nil, err
	}

	forLen := len(articles)
	// 封装查询结果为[]*content.SelectArticles类型
	respArticles := make([]*content.SelectArticles, forLen)

	for i := 0; i < forLen; i++ {
		currentArticle := articles[i]

		// tags属性中原本包含的是tagid, 修改article中的tagid为tagname, 即根据id查询tag再获取tagName
		selectTagByIdLogic := NewSelectTagByIdLogic(l.ctx, l.svcCtx)
		for i := 0; i < len(currentArticle.Tags); i++ {
			tag, _ := selectTagByIdLogic.SelectTagById(&content.SelectTagByIdRequest{
				Id: currentArticle.Tags[i],
			})

			currentArticle.Tags[i] = tag.Name
		}

		respArticles[i] = &content.SelectArticles{
			Id:    currentArticle.ID.Hex(),
			Title: currentArticle.Title,
			Desc:  currentArticle.Desc,
			Tags:  currentArticle.Tags,
		}
	}

	return &content.SelectArticleByPageResponse{
		Articles: respArticles,
		Total:    total,
	}, nil
}
