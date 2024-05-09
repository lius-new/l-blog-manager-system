package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/user/internal/svc"
	"github.com/lius-new/blog-backend/api/user/internal/types"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyStatusLogic {
	return &ModifyStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyStatusLogic) ModifyStatus(req *types.ModifyStatusRequest) (resp *types.CreateResponse, err error) {
	_, err = l.svcCtx.Userer.ModifyStatus(l.ctx, &user.ModifyUserStatusRequest{
		Id:     req.Id,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{}, nil
}
