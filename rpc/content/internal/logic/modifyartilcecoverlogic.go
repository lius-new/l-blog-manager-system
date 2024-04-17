package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/article"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyArtilceCoverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArtilceCoverLogic {
	return &ModifyArtilceCoverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改文章Cover
func (l *ModifyArtilceCoverLogic) ModifyArtilceCover(in *content.ModifyArticleCoverRequest) (*content.ModifyArticleCoverResponse, error) {
	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, rpc.ErrInvalidObjectId
	}
	_, err = l.svcCtx.ModelWithArticle.Update(l.ctx, &model.Article{
		ID:     id,
		Covers: in.Covers,
	})
	if err != nil {
		return nil, err
	}

	return &content.ModifyArticleCoverResponse{}, nil
}
