package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectCoverByHashLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectCoverByHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectCoverByHashLogic {
	return &SelectCoverByHashLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询cover
func (l *SelectCoverByHashLogic) SelectCoverByHash(in *content.SelectCoverByHashRequest) (*content.SelectCoverByHashResponse, error) {
	if len(in.Hash) == 0 {
		return nil, rpc.ErrRequestParam
	}

	currentCover, err := l.svcCtx.ModelWithCover.FindOneByHash(l.ctx, in.Hash)

	if err == rpc.ErrNotFound || currentCover == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &content.SelectCoverByHashResponse{
		Cover: &content.Cover{
			Id:      currentCover.ID.Hex(),
			Content: currentCover.Content,
			Hash:    currentCover.Hash,
		},
	}, nil
}
