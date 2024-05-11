package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type ExistTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExistTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExistTagLogic {
	return &ExistTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// * tag exist *
func (l *ExistTagLogic) ExistTag(in *content.ExistTagRequest) (*content.ExistTagResponse, error) {
	findTag, err := l.svcCtx.ModelWithTag.FindOne(l.ctx, in.Id)

	if err == rpc.ErrNotFound || findTag == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &content.ExistTagResponse{
		Exist: true,
	}, nil
}
