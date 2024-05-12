package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectWhiteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectWhiteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectWhiteListLogic {
	return &SelectWhiteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectWhiteListLogic) SelectWhiteList(in *analyzer.SelectWhiteListRequest) (*analyzer.SelectWhiteListResponse, error) {
	// todo: add your logic here and delete this line

	return &analyzer.SelectWhiteListResponse{}, nil
}
