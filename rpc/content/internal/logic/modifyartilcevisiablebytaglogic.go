package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type ModifyArtilceVisiableByTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceVisiableByTagLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *ModifyArtilceVisiableByTagLogic {
	return &ModifyArtilceVisiableByTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据tag修改文章的可见性
func (l *ModifyArtilceVisiableByTagLogic) ModifyArtilceVisiableByTag(
	in *content.ModifyArticleVisiableByTagRequest,
) (*content.ModifyArticleVisiableByTagResponse, error) {

	if len(in.TagId) == 0 {
		return nil, rpc.ErrRequestParam
	}

	// 判断文章是否存在
	currentTag, err := l.svcCtx.ModelWithTag.FindOne(l.ctx, in.TagId)
	if err == rpc.ErrNotFound || currentTag == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// 遍历指定tag下所有的article
	modifyArtilceVisiableLogic := NewModifyArtilceVisiableLogic(l.ctx, l.svcCtx)
	for _, v := range currentTag.Articles {
		_, err := modifyArtilceVisiableLogic.ModifyArtilceVisiable(&content.ModifyArticleVisiableRequest{
			Id:       v,
			Visiable: in.Visiable,
		})
		// TODO: 感觉忽略错误更好
		if err != nil {
			return nil, err
		}
	}

	return &content.ModifyArticleVisiableByTagResponse{}, nil
}
