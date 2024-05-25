package logic

import (
	"context"
	"time"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyBlockedWithBlockEndLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyBlockedWithBlockEndLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyBlockedWithBlockEndLogic {
	return &ModifyBlockedWithBlockEndLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyBlockedWithBlockEndLogic) ModifyBlockedWithBlockEnd(in *analyzer.ModifyBlockedWithBlockEndRequest) (*analyzer.ModifyBlockedWithBlockEndResponse, error) {
	// 先查询到指定的blocked
	selectBlockedByBlockIPLogic := NewSelectBlockedByBlockIPLogic(l.ctx, l.svcCtx)
	selectResp, err := selectBlockedByBlockIPLogic.SelectBlockedByBlockIP(&analyzer.SelectBlockedByBlockIPRequest{
		BlockIp: in.BlockIp,
	})

	if err != nil {
		return nil, err
	}

	// 以秒为单位
	endTime := time.Unix(selectResp.BlockEnd+in.BlockEnd, 0)

	// 在指定的blocked上加一
	err = l.svcCtx.ModelWithBlocked.ModifyBlockByBlockIPWithBlockend(l.ctx, in.BlockIp, endTime)
	if err != nil {
		return nil, err
	}

	return &analyzer.ModifyBlockedWithBlockEndResponse{}, nil
}
