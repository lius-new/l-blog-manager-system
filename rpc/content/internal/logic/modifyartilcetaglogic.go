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

type ModifyArtilceTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArtilceTagLogic {
	return &ModifyArtilceTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改文章标签
func (l *ModifyArtilceTagLogic) ModifyArtilceTag(in *content.ModifyArticleTagRequest) (*content.ModifyArticleTagResponse, error) {
	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, rpc.ErrInvalidObjectId
	}
	_, err = l.svcCtx.ModelWithArticle.Update(l.ctx, &model.Article{
		ID:   id,
		Tags: in.Tags,
	})
	if err != nil {
		return nil, err
	}
	return &content.ModifyArticleTagResponse{}, nil
}