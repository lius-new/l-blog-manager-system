package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/user/internal/svc"
	"github.com/lius-new/blog-backend/api/user/internal/types"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyUserNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyUserNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyUserNameLogic {
	return &ModifyUserNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyUserNameLogic) ModifyUserName(req *types.ModifyUserNameRequest) (resp *types.CreateResponse, err error) {
	_, err = l.svcCtx.Userer.ModifyUserName(l.ctx, &user.ModifyUserNameRequest{
		Id:       req.Id,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{}, nil
}
