package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type SelectCoverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectCoverLogic {
	return &SelectCoverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询cover
func (l *SelectCoverLogic) SelectCover(
	in *content.SelectCoverRequest,
) (*content.SelectCoverResponse, error) {
	if len(in.Id) == 0 {
		return nil, rpc.ErrRequestParam
	}

	currentCover, err := l.svcCtx.ModelWithCover.FindOne(l.ctx, in.Id)

	if err == rpc.ErrNotFound || currentCover == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &content.SelectCoverResponse{
		Cover: &content.Cover{
			Id:      currentCover.ID.Hex(),
			Content: currentCover.Content,
			Hash:    currentCover.Hash,
		},
	}, nil
}
