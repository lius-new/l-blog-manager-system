package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectBlockedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectBlockedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectBlockedLogic {
	return &SelectBlockedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectBlockedLogic) SelectBlocked(in *analyzer.SelectBlockedRequest) (*analyzer.SelectBlockedResponse, error) {
	// todo: add your logic here and delete this line

	return &analyzer.SelectBlockedResponse{}, nil
}
