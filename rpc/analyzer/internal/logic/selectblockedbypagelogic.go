package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectBlockedByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectBlockedByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectBlockedByPageLogic {
	return &SelectBlockedByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectBlockedByPageLogic) SelectBlockedByPage(in *analyzer.SelectBlockedByPageRequest) (*analyzer.SelectBlockedByPageResponse, error) {
	// todo: add your logic here and delete this line

	return &analyzer.SelectBlockedByPageResponse{}, nil
}
