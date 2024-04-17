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

type ModifyArtilceContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyArtilceContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyArtilceContentLogic {
	return &ModifyArtilceContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改文章内容
func (l *ModifyArtilceContentLogic) ModifyArtilceContent(in *content.ModifyArticleContentRequest) (*content.ModifyArticleContentResponse, error) {

	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, rpc.ErrInvalidObjectId
	}
	_, err = l.svcCtx.ModelWithArticle.Update(l.ctx, &model.Article{
		ID:      id,
		Content: in.Content,
	})
	if err != nil {
		return nil, err
	}
	return &content.ModifyArticleContentResponse{}, nil
}
