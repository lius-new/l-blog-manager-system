package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyWhiteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyWhiteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyWhiteListLogic {
	return &ModifyWhiteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyWhiteListLogic) ModifyWhiteList(in *analyzer.ModifyWhiteListRequest) (*analyzer.ModifyWhiteListResponse, error) {
	// todo: add your logic here and delete this line

	return &analyzer.ModifyWhiteListResponse{}, nil
}
