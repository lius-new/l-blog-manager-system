package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectWhiteListByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectWhiteListByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectWhiteListByPageLogic {
	return &SelectWhiteListByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectWhiteListByPageLogic) SelectWhiteListByPage(in *analyzer.SelectWhiteListByPageRequest) (*analyzer.SelectWhiteListByPageResponse, error) {
	// todo: add your logic here and delete this line

	return &analyzer.SelectWhiteListByPageResponse{}, nil
}
