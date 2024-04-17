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

type ModifyTagVisiableLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyTagVisiableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyTagVisiableLogic {
	return &ModifyTagVisiableLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改tag可见性(visiable)
func (l *ModifyTagVisiableLogic) ModifyTagVisiable(in *content.ModifyTagVisiableRequest) (*content.ModifyTagVisiableResponse, error) {
	id, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, rpc.ErrInvalidObjectId
	}

	_, err = l.svcCtx.ModelWithTag.Update(l.ctx, &model.Tag{
		ID:       id,
		Visiable: in.Visiable,
	})
	if err != nil {
		return nil, err
	}
	return &content.ModifyTagVisiableResponse{}, nil
}
