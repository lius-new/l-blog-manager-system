package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyBlockedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyBlockedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyBlockedLogic {
	return &ModifyBlockedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyBlockedLogic) ModifyBlocked(in *analyzer.ModifyBlockedRequest) (*analyzer.ModifyBlockedResponse, error) {
	// todo: add your logic here and delete this line

	return &analyzer.ModifyBlockedResponse{}, nil
}
