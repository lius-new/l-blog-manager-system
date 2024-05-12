package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBlockedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBlockedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBlockedLogic {
	return &CreateBlockedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ================ Blocked  ================
func (l *CreateBlockedLogic) CreateBlocked(in *analyzer.CreateBlockedRequest) (*analyzer.CreateBlockedResponse, error) {
	// todo: add your logic here and delete this line

	return &analyzer.CreateBlockedResponse{}, nil
}
