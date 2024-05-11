package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/tag"
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
func (l *ModifyTagNameLogic) ModifyTagName(
	in *content.ModifyTagNameRequest,
) (*content.ModifyTagNameResponse, error) {
	// 判断指定tag是否存在
	if _, err := NewExistTagLogic(l.ctx, l.svcCtx).ExistTag(&content.ExistTagRequest{
		Id: in.Id,
	}); err != nil {
		return nil, err
	}

	// 封装id
	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, rpc.ErrInvalidObjectId
	}
	// 更新
	_, err = l.svcCtx.ModelWithTag.Update(l.ctx, &model.Tag{
		ID:   id,
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}

	return &content.ModifyTagNameResponse{}, nil
}
