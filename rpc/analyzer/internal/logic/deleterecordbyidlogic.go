package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRecordByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRecordByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRecordByIdLogic {
	return &DeleteRecordByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRecordByIdLogic) DeleteRecordById(in *analyzer.DeleteRecordByIdRequest) (*analyzer.DeleteRecordByIdResponse, error) {
	_, err := l.svcCtx.ModelWithRecord.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &analyzer.DeleteRecordByIdResponse{}, nil
}
