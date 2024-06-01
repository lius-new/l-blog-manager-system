package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/article"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArtilceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArtilceLogic {
	return &ModifyArtilceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyArtilceLogic) ModifyArtilce(in *content.ModifyArticleRequest) (*content.ModifyArticleResponse, error) {
	if len(in.Title) == 0 || len(in.Description) == 0 || len(in.Content) == 0 {
		return nil, rpc.ErrRequestParam
	}

	// 判断文章是否存在
	currentArticle, err := l.svcCtx.ModelWithArticle.FindOne(l.ctx, in.Id)
	if err == rpc.ErrNotFound || currentArticle == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// 更新cover
	// 删除原本的图片
	for _, v := range currentArticle.Covers {
		NewDeleteCoverLogic(l.ctx, l.svcCtx).DeleteCover(&content.DeleteCoverRequest{
			Id: v,
		})
	}
	// 保存图片
	coverIds, _ := NewCreateCoversLogic(
		l.ctx,
		l.svcCtx,
	).CreateCovers(&content.CreateCoversRequest{
		Content: in.Covers,
	})

	// TODO: 搞不懂为什么结果体只设置指定属性那么其他属性就会设置为对应零值，日志也显示只修改了指定属性而没有修改其他属性呀
	// 更新
	_, err = l.svcCtx.ModelWithArticle.Update(l.ctx, &model.Article{
		ID:       currentArticle.ID,
		Title:    in.Title,
		Desc:     in.Description,
		Content:  in.Content,
		Tags:     currentArticle.Tags,
		Covers:   coverIds.Ids,
		Visiable: in.Visiable,
	})
	if err != nil {
		return nil, err
	}

	// tag 调用 NewModifyArtilceTagLogic更新, 因为 更新tag逻辑实在是有些长
	_, err = NewModifyArtilceTagLogic(l.ctx, l.svcCtx).ModifyArtilceTag(&content.ModifyArticleTagRequest{
		Id:   currentArticle.ID.Hex(),
		Tags: in.Tags,
	})
	if err != nil {
		return nil, err
	}

	return &content.ModifyArticleResponse{}, nil
}
