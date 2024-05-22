package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/article"
)

type ModifyArtilceDescLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceDescLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *ModifyArtilceDescLogic {
	return &ModifyArtilceDescLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改文章描述
func (l *ModifyArtilceDescLogic) ModifyArtilceDesc(
	in *content.ModifyArticleDescRequest,
) (*content.ModifyArticleDescResponse, error) {
	if len(in.Id) == 0 || len(in.Desc) == 0 {
		return nil, rpc.ErrRequestParam
	}
	// 判断文章是否存在
	currentArticle, err := l.svcCtx.ModelWithArticle.FindOne(l.ctx, in.Id)
	if err == rpc.ErrNotFound || currentArticle == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// TODO: 搞不懂为什么结果体只设置指定属性那么其他属性就会设置为对应零值，日志也显示只修改了指定属性而没有修改其他属性呀
	// 更新
	_, err = l.svcCtx.ModelWithArticle.Update(l.ctx, &model.Article{
		ID:       currentArticle.ID,
		Title:    currentArticle.Title,
		Desc:     in.Desc,
		Content:  currentArticle.Content,
		Tags:     currentArticle.Tags,
		Covers:   currentArticle.Covers,
		Visiable: currentArticle.Visiable,
	})
	if err != nil {
		return nil, err
	}

	return &content.ModifyArticleDescResponse{}, nil
}
