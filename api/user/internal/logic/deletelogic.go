package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/user/user"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/api/user/internal/svc"
	"github.com/lius-new/blog-backend/api/user/internal/types"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.IdRequest) (resp *types.DeleteResponse, err error) {
	_, err = l.svcCtx.Userer.Delete(l.ctx, &user.DeleteUserRequest{
		Uid: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteResponse{}, nil
}
