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

type ModifyArtilceVisiableLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceVisiableLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *ModifyArtilceVisiableLogic {
	return &ModifyArtilceVisiableLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改文章的可见性
func (l *ModifyArtilceVisiableLogic) ModifyArtilceVisiable(
	in *content.ModifyArticleVisiableRequest,
) (*content.ModifyArticleVisiableResponse, error) {
	if len(in.Id) == 0 {
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
		ID:       id,
		Visiable: in.Visiable,
	})
	if err != nil {
		return nil, err
	}
	return &content.ModifyArticleVisiableResponse{}, nil
}
