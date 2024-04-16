package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/user/internal/svc"
	"github.com/lius-new/blog-backend/api/user/internal/types"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
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
	deleteResp, err := l.svcCtx.Userer.Delete(l.ctx, &user.DeleteUserRequest{
		Uid: req.Id,
	})

	if err != nil {
		return &types.DeleteResponse{
			Id:     deleteResp.Uid,
			Status: false,
		}, err
	}

	return &types.DeleteResponse{
		Id:     deleteResp.Uid,
		Status: false,
	}, nil
}
