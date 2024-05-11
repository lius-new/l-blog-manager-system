package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type DeleteCoverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCoverLogic {
	return &DeleteCoverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除cover
func (l *DeleteCoverLogic) DeleteCover(
	in *content.DeleteCoverRequest,
) (*content.DeleteCoverResponse, error) {
	count, err := l.svcCtx.ModelWithCover.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &content.DeleteCoverResponse{Count: count}, nil
}
