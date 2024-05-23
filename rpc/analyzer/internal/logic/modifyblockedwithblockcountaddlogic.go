package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyBlockedWithBlockCountAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyBlockedWithBlockCountAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyBlockedWithBlockCountAddLogic {
	return &ModifyBlockedWithBlockCountAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyBlockedWithBlockCountAddLogic) ModifyBlockedWithBlockCountAdd(in *analyzer.ModifyBlockedWithBlockCountAddRequest) (*analyzer.ModifyBlockedWithBlockCountAddResponse, error) {
	// 先查询到指定的blocked
	selectBlockedByBlockIPLogic := NewSelectBlockedByBlockIPLogic(l.ctx, l.svcCtx)
	selectResp, err := selectBlockedByBlockIPLogic.SelectBlockedByBlockIP(&analyzer.SelectBlockedByBlockIPRequest{
		BlockIp: in.BlockIp,
	})

	if err != nil {
		return nil, err
	}

	// 在指定的blocked上加一
	err = l.svcCtx.ModelWithBlocked.ModifyBlockByBlockIPWithCount(l.ctx, in.BlockIp, selectResp.BlockCount+1)
	if err != nil {
		return nil, err
	}

	return &analyzer.ModifyBlockedWithBlockCountAddResponse{}, nil
}
