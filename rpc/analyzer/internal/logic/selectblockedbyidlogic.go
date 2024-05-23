package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectBlockedByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectBlockedByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectBlockedByIdLogic {
	return &SelectBlockedByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectBlockedByIdLogic) SelectBlockedById(in *analyzer.SelectBlockedByIdRequest) (*analyzer.SelectBlockedByIdResponse, error) {
	blocked, err := l.svcCtx.ModelWithBlocked.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &analyzer.SelectBlockedByIdResponse{
		Id:         blocked.ID.Hex(),
		BlockIP:    blocked.BlockIP,
		BlockEnd:   blocked.BlockEnd.Unix(),
		BlockCount: blocked.BlockCount,
		CreateAt:   blocked.CreateAt.Unix(),
		UpdateAt:   blocked.UpdateAt.Unix(),
	}, nil
}
