package logic

import (
	"context"
	"time"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/analyzer/model/mongo/record"

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

// MergeRecord: 每次合并表示封禁
func (l *MergeRecordLogic) MergeRecord(in *analyzer.MergeRecordRequest) (*analyzer.MergeRecordResponse, error) {
	// 查询系统设置，采用最新的系统设置(最后一条)。
	setting, err := l.svcCtx.ModelWithSetting.FindLastSetting(l.ctx)
	if err != nil && err != rpc.ErrNotFound {
		return nil, err
	}

	// 如果没有添加设置，那么设置默认的合并边界和合并(时间)范围
	if err == rpc.ErrNotFound {
		setting.RecordMergeBoundary = 50
		setting.RecordMergeInterval = time.Hour * 12
	}

	recordInternalCount, err := l.svcCtx.ModelWithRecord.CountScopeTimeRecordNumber(l.ctx, in.RequestIp, setting.RecordMergeInterval)
	if err != nil {
		return nil, err
	}

	// 不需要和合并
	if recordInternalCount <= setting.RecordMergeBoundary {
		return nil, nil
	}

	// 后续是合并操作
	// 1. 删除指定时间范围内指定IP的日志
	deleteCount, err := l.svcCtx.ModelWithRecord.DeleteScopeTimeRecord(l.ctx, in.RequestIp, setting.RecordMergeInterval)
	if err != nil {
		return nil, err
	}
	// 2. 插入一条新的数据, 为合并后的数据
	err = l.svcCtx.ModelWithRecord.Insert(l.ctx, &model.Record{
		RequestIP:     in.RequestIp,
		RequestMethod: "*",
		RequestPath:   "*",
		RequestCount:  deleteCount,
	})
	if err != nil {
		return nil, err
	}

	// 最后调用封禁
	_, err = NewCreateBlockedLogic(l.ctx, l.svcCtx).CreateBlocked(&analyzer.CreateBlockedRequest{
		BlockIp: in.RequestIp,
	})
	if err != nil {
		return nil, err
	}

	return &analyzer.MergeRecordResponse{}, nil
}
