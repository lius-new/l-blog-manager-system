package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/analyzer/analyzer"
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectBlockedByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectBlockedByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectBlockedByPageLogic {
	return &SelectBlockedByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectBlockedByPageLogic) SelectBlockedByPage(in *analyzer.SelectBlockedByPageRequest) (*analyzer.SelectBlockedByPageResponse, error) {
	blockeds, total, err := l.svcCtx.ModelWithBlocked.FindByPage(l.ctx, in.PageNum, in.PageSize)
	if err != nil {
		return nil, err
	}

	// 封装数据
	forLen := len(blockeds)
	data := make([]*analyzer.SelectBlockeds, forLen)
	for i := 0; i < forLen; i++ {
		currentBlocked := blockeds[i]
		data[i] = &analyzer.SelectBlockeds{
			Id:         currentBlocked.ID.Hex(),
			BlockIP:    currentBlocked.BlockIP,
			BlockEnd:   currentBlocked.BlockEnd.Unix(),
			BlockCount: currentBlocked.BlockCount,
			CreateAt:   currentBlocked.CreateAt.Unix(),
			UpdateAt:   currentBlocked.UpdateAt.Unix(),
		}
	}

	return &analyzer.SelectBlockedByPageResponse{
		Blockeds: data,
		Total:    total,
	}, nil
}
