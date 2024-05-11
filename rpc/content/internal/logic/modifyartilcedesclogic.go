package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"

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
	if _, err := NewExistArtilceLogic(l.ctx, l.svcCtx).ExistArtilce(&content.ExistArtilceRequest{
		Id: in.Id,
	}); err != nil {
		return nil, err
	}
	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, rpc.ErrInvalidObjectId
	}
	_, err = l.svcCtx.ModelWithArticle.Update(l.ctx, &model.Article{
		ID:   id,
		Desc: in.Desc,
	})
	if err != nil {
		return nil, err
	}

	return &content.ModifyArticleDescResponse{}, nil
}
