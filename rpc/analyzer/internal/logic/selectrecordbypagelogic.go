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
	// todo: add your logic here and delete this line

	return &analyzer.SelectRecordByPageResponse{}, nil
}
