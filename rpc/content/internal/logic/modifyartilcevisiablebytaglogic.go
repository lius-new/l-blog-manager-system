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

type ModifyArtilceVisiableByTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceVisiableByTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArtilceVisiableByTagLogic {
	return &ModifyArtilceVisiableByTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据tag修改文章的可见性
func (l *ModifyArtilceVisiableByTagLogic) ModifyArtilceVisiableByTag(in *content.ModifyArticleVisiableByTagRequest) (*content.ModifyArticleVisiableByTagResponse, error) {
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
	return &content.ModifyArticleVisiableByTagResponse{}, nil
}
