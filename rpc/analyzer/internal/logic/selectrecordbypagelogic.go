package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectRecordByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectRecordByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectRecordByPageLogic {
	return &SelectRecordByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectRecordByPageLogic) SelectRecordByPage(in *analyzer.SelectRecordByPageRequest) (*analyzer.SelectRecordByPageResponse, error) {

	records, total, err := l.svcCtx.ModelWithRecord.FindByPage(l.ctx, in.PageNum, in.PageSize)
	if err != nil {
		return nil, err
	}

	// 封装数据
	forLen := len(records)
	data := make([]*analyzer.SelectRecords, forLen)
	for i := 0; i < forLen; i++ {
		currentRecord := records[i]
		data[i] = &analyzer.SelectRecords{
			Id:            currentRecord.ID.Hex(),
			RequestIP:     currentRecord.RequestIP,
			RequestMethod: currentRecord.RequestMethod,
			RequestPath:   currentRecord.RequestPath,
			RequestCount:  currentRecord.RequestCount,
			CreateAt:      currentRecord.CreateAt.Unix(),
			UpdateAt:      currentRecord.UpdateAt.Unix(),
		}
	}

	return &analyzer.SelectRecordByPageResponse{
		Records: data,
		Total:   total,
	}, nil
}
