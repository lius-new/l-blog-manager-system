package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWhiteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateWhiteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWhiteListLogic {
	return &CreateWhiteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ================ whitelist ================
func (l *CreateWhiteListLogic) CreateWhiteList(in *analyzer.CreateWhiteListRequest) (*analyzer.CreateWhiteListResponse, error) {
	// todo: add your logic here and delete this line

	return &analyzer.CreateWhiteListResponse{}, nil
}
