package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

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
	// todo: add your logic here and delete this line

	return &content.ModifyTagNameResponse{}, nil
}
