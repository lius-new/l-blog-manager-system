package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/tag"
)

type CreateTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTagLogic {
	return &CreateTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ** tag **
func (l *CreateTagLogic) CreateTag(
	in *content.CreateTagRequest,
) (*content.CreateTagResponse, error) {
	t, err := l.svcCtx.ModelWithTag.FindByName(l.ctx, in.Name)

	// 如果t不等nil就意味tag已经存在
	if t != nil {
		return &content.CreateTagResponse{
			Id: t.ID.Hex(),
		}, nil
	}
	// 如果存在错误且错误不是NotFound那么就抛出
	if err != nil && err != rpc.ErrNotFound {
		return nil, err
	}

	id, err := l.svcCtx.ModelWithTag.InsertReturnId(l.ctx, &model.Tag{
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}

	return &content.CreateTagResponse{
		Id: id,
	}, nil
}
