package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/user/internal/svc"
	"github.com/lius-new/blog-backend/api/user/internal/types"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyPasswordLogic {
	return &ModifyPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyPasswordLogic) ModifyPassword(req *types.ModifyPasswordRequest) (resp *types.CreateResponse, err error) {
	_, err = l.svcCtx.Userer.ModifyPassword(l.ctx, &user.ModifyPasswordRequest{
		Id:       req.Id,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{}, nil
}
