package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/tag"

	"github.com/zeromicro/go-zero/core/logx"
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
func (l *CreateTagLogic) CreateTag(in *content.CreateTagRequest) (*content.CreateTagResponse, error) {
	err := l.svcCtx.ModelWithTag.Insert(l.ctx, &model.Tag{
		Name:     in.Name,
		Articles: in.Articles,
		Visiable: in.Visiable,
	})
	if err != nil {
		return nil, err
	}

	return &content.CreateTagResponse{}, nil
}
