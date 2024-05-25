package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/analyzer/model/mongo/record"

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
	if len(in.RequestIp) == 0 || len(in.RequestMethod) == 0 || len(in.RequestPath) == 0 {
		return nil, rpc.ErrRequestParam
	}

	err := l.svcCtx.ModelWithRecord.Insert(l.ctx, &model.Record{
		RequestIP:     in.RequestIp,
		RequestMethod: in.RequestMethod,
		RequestPath:   in.RequestPath,
		RequestCount:  1,
	})

	if err != nil {
		return nil, err
	}

	// 检查是否需要合并日志，（合并会伴随block）
	_, err = NewMergeRecordLogic(l.ctx, l.svcCtx).MergeRecord(&analyzer.MergeRecordRequest{
		RequestIp: in.RequestIp,
	})
	if err != nil {
		return nil, err
	}

	return &analyzer.CreateRecordResponse{}, nil
}
