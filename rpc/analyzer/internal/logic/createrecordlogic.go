package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRecordLogic {
	return &CreateRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ================ Record  ================
func (l *CreateRecordLogic) CreateRecord(in *analyzer.CreateRecordRequest) (*analyzer.CreateRecordResponse, error) {
	// todo: add your logic here and delete this line

	return &analyzer.CreateRecordResponse{}, nil
}
