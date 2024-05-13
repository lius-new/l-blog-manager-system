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

type ModifyArtilceTitleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceTitleLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *ModifyArtilceTitleLogic {
	return &ModifyArtilceTitleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// * article modify *
func (l *ModifyArtilceTitleLogic) ModifyArtilceTitle(
	in *content.ModifyArticleTitleRequest,
) (*content.ModifyArticleTitleResponse, error) {
  if len(in.Id) == 0 || len(in.Title) == 0 {
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
		ID:    id,
		Title: in.Title,
	})
	if err != nil {
		return nil, err
	}

	return &content.ModifyArticleTitleResponse{}, nil
}
