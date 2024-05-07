package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/user/internal/svc"
	"github.com/lius-new/blog-backend/api/user/internal/types"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyLogic {
	return &ModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyLogic) Modify(req *types.ModifyRequest) (resp *types.CreateResponse, err error) {
	modifyResp, err := l.svcCtx.Userer.Modify(l.ctx, &user.ModifyUserRequest{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Status:   req.Status,
	})
	if err != nil {
		return &types.CreateResponse{
			Status: false,
		}, err
	}

	return &types.CreateResponse{
		Data: types.UserBackend{
			Id:       req.Id,
			Username: modifyResp.Username,
			Status:   modifyResp.Status,
		},
		Status: true,
	}, nil
}
