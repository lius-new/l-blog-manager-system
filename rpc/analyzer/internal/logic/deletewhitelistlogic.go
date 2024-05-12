package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteWhiteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteWhiteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteWhiteListLogic {
	return &DeleteWhiteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteWhiteListLogic) DeleteWhiteList(in *analyzer.DeleteWhiteListRequest) (*analyzer.DeleteWhiteListResponse, error) {
	// todo: add your logic here and delete this line

	return &analyzer.DeleteWhiteListResponse{}, nil
}
