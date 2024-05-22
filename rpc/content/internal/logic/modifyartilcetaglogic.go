package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/article"
)

type ModifyArtilceTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceTagLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *ModifyArtilceTagLogic {
	return &ModifyArtilceTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改文章标签
func (l *ModifyArtilceTagLogic) ModifyArtilceTag(
	in *content.ModifyArticleTagRequest,
) (*content.ModifyArticleTagResponse, error) {
	if len(in.Id) == 0 || len(in.Tags) == 0 {
		return nil, rpc.ErrRequestParam
	}

	// 判断文章是否存在
	currentArticle, err := l.svcCtx.ModelWithArticle.FindOne(l.ctx, in.Id)
	if err == rpc.ErrNotFound || currentArticle == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// 过滤空字符串
	names := make([]string, 0)
	for _, v := range in.Tags {
		if len(v) > 0 {
			names = append(names, v)
		}
	}

	// 保存tag
	tagIds, _ := NewCreateTagsLogic(
		l.ctx,
		l.svcCtx,
	).CreateTags(&content.CreateTagsRequest{
		Names: names,
	})

	// 遍历之前的tag将指定article移除
	modifyTagRemoveArticleLogic := NewModifyTagRemoveArticleLogic(l.ctx, l.svcCtx)
	for _, tagId := range currentArticle.Tags {
		modifyTagRemoveArticleLogic.ModifyTagRemoveArticle(&content.ModifyTagRemoveArticleRequest{
			Id:      tagId,
			Article: currentArticle.ID.Hex(),
		})
	}

	// 将指定文章添加到tag
	modifyTagPushArticleLogic := NewModifyTagPushArticleLogic(l.ctx, l.svcCtx)
	for _, v := range tagIds.Ids {
		modifyTagPushArticleLogic.ModifyTagPushArticle(&content.ModifyTagPushArticleRequest{
			Id:      v,
			Article: currentArticle.ID.Hex(),
		})
	}

	// TODO: 搞不懂为什么结果体只设置指定属性那么其他属性就会设置为对应零值，日志也显示只修改了指定属性而没有修改其他属性呀
	// 更新
	_, err = l.svcCtx.ModelWithArticle.Update(l.ctx, &model.Article{
		ID:       currentArticle.ID,
		Title:    currentArticle.Title,
		Desc:     currentArticle.Desc,
		Content:  currentArticle.Content,
		Tags:     tagIds.Ids,
		Covers:   currentArticle.Covers,
		Visiable: currentArticle.Visiable,
	})
	if err != nil {
		return nil, err
	}
	return &content.ModifyArticleTagResponse{}, nil
}
