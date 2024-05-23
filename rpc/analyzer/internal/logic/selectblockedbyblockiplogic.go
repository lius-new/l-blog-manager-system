package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectBlockedByBlockIPLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectBlockedByBlockIPLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectBlockedByBlockIPLogic {
	return &SelectBlockedByBlockIPLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectBlockedByBlockIPLogic) SelectBlockedByBlockIP(in *analyzer.SelectBlockedByBlockIPRequest) (*analyzer.SelectBlockedByBlockIPResponse, error) {
	blocked, err := l.svcCtx.ModelWithBlocked.FindByBlockIP(l.ctx, in.BlockIp)

	if err != nil {
		return nil, err
	}

	return &analyzer.SelectBlockedByBlockIPResponse{
		BlockIP:    blocked.BlockIP,
		BlockEnd:   blocked.BlockEnd.Unix(),
		BlockCount: blocked.BlockCount,
		CreateAt:   blocked.CreateAt.Unix(),
		UpdateAt:   blocked.UpdateAt.Unix(),
	}, nil
}
