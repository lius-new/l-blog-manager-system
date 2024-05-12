package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MergeRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMergeRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergeRecordLogic {
	return &MergeRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MergeRecordLogic) MergeRecord(in *analyzer.MergeRecordRequest) (*analyzer.MergeRecordResponse, error) {
	// todo: add your logic here and delete this line

	return &analyzer.MergeRecordResponse{}, nil
}
