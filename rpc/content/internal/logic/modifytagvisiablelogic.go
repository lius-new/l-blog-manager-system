package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

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
	// todo: add your logic here and delete this line

	return &content.ModifyTagVisiableResponse{}, nil
}
