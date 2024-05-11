package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/tag"
)

type ModifyTagVisiableLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyTagVisiableLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *ModifyTagVisiableLogic {
	return &ModifyTagVisiableLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改tag可见性(visiable)
func (l *ModifyTagVisiableLogic) ModifyTagVisiable(
	in *content.ModifyTagVisiableRequest,
) (*content.ModifyTagVisiableResponse, error) {
	// 判断指定tag是否存在

	currentTag, err := l.svcCtx.ModelWithTag.FindOne(l.ctx, in.Id)
	if err == rpc.ErrNotFound || currentTag == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.ModelWithTag.Update(l.ctx, &model.Tag{
		ID:       currentTag.ID,
		Visiable: in.Visiable,
	})
	if err != nil {
		return nil, err
	}
	return &content.ModifyTagVisiableResponse{}, nil
}
