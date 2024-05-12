package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBlockedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBlockedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBlockedLogic {
	return &DeleteBlockedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteBlockedLogic) DeleteBlocked(in *analyzer.DeleteBlockedRequest) (*analyzer.DeleteBlockedResponse, error) {
	// todo: add your logic here and delete this line

	return &analyzer.DeleteBlockedResponse{}, nil
}
