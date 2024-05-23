package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectRecordByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectRecordByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectRecordByIdLogic {
	return &SelectRecordByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectRecordByIdLogic) SelectRecordById(in *analyzer.SelectRecordByIdRequest) (*analyzer.SelectRecordByIdResponse, error) {
	record, err := l.svcCtx.ModelWithRecord.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &analyzer.SelectRecordByIdResponse{
		Id:            record.ID.Hex(),
		RequestIP:     record.RequestIP,
		RequestMethod: record.RequestMethod,
		RequestPath:   record.RequestPath,
		RequestCount:  record.RequestCount,
		CreateAt:      record.CreateAt.Unix(),
		UpdateAt:      record.UpdateAt.Unix(),
	}, nil
}
