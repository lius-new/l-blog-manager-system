package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyRecordLogic {
	return &ModifyRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyRecordLogic) ModifyRecord(in *analyzer.ModifyRecordRequest) (*analyzer.ModifyRecordResponse, error) {
	// TODO: 考虑这个接口是否应该实现

	return &analyzer.ModifyRecordResponse{}, nil
}
