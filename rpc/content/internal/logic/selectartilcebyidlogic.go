package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type SelectArtilceByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectArtilceByIdLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *SelectArtilceByIdLogic {
	return &SelectArtilceByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// * article select *
func (l *SelectArtilceByIdLogic) SelectArtilceById(
	in *content.SelectArticleByIdRequest,
) (*content.SelectArticle, error) {

	// 查询指定id的文章
	findArticle, err := l.svcCtx.ModelWithArticle.FindOne(l.ctx, in.Id)

	if err == rpc.ErrNotFound || findArticle == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	// tags属性中原本包含的是tagid, 修改article中的tagid为tagname, 即根据id查询tag再获取tagName
	selectTagByIdLogic := NewSelectTagByIdLogic(l.ctx, l.svcCtx)
	for i := 0; i < len(findArticle.Tags); i++ {
		tag, _ := selectTagByIdLogic.SelectTagById(&content.SelectTagByIdRequest{
			Id: findArticle.Tags[i],
		})

		findArticle.Tags[i] = tag.Name
	}
	// covers属性中原本包含的是coverId, 修改article中的coverId为cover hash, 即根据id查询cover再获取hash
	selectCoverLogic := NewSelectCoverLogic(l.ctx, l.svcCtx)
	for i := 0; i < len(findArticle.Covers); i++ {
		cover, _ := selectCoverLogic.SelectCover(&content.SelectCoverRequest{
			Id: findArticle.Covers[i],
		})
		findArticle.Covers[i] = cover.Cover.Hash
	}

	return &content.SelectArticle{
		Id:      findArticle.ID.Hex(),
		Title:   findArticle.Title,
		Desc:    findArticle.Desc,
		Content: findArticle.Content,
		Tags:    findArticle.Tags,
		Covers:  findArticle.Covers,
		Time:    findArticle.UpdateAt.Unix(),
	}, nil
}
