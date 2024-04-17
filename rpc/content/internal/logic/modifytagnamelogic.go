package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/tag"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyTagNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyTagNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyTagNameLogic {
	return &ModifyTagNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改tag name
func (l *ModifyTagNameLogic) ModifyTagName(in *content.ModifyTagNameRequest) (*content.ModifyTagNameResponse, error) {
	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, rpc.ErrInvalidObjectId
	}

	_, err = l.svcCtx.ModelWithTag.Update(l.ctx, &model.Tag{
		ID:   id,
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}

	return &content.ModifyTagNameResponse{}, nil
}
