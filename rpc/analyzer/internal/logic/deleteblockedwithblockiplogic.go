package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBlockedWithBlockIPLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBlockedWithBlockIPLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBlockedWithBlockIPLogic {
	return &DeleteBlockedWithBlockIPLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据指定ip删除BLCOKED (实际上只是设置blocked 时间归零)
func (l *DeleteBlockedWithBlockIPLogic) DeleteBlockedWithBlockIP(in *analyzer.DeleteBlockedWithBlockIPRequest) (*analyzer.DeleteBlockedWithBlockIPResponse, error) {
	_, err := l.svcCtx.ModelWithBlocked.DeleteBlockByBlockIP(l.ctx, in.BlockIp)
	if err != nil {
		return nil, err
	}

	return &analyzer.DeleteBlockedWithBlockIPResponse{}, nil
}
