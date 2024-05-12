package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectRecordLogic {
	return &SelectRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectRecordLogic) SelectRecord(in *analyzer.SelectRecordRequest) (*analyzer.SelectRecordResponse, error) {
	// todo: add your logic here and delete this line

	return &analyzer.SelectRecordResponse{}, nil
}
