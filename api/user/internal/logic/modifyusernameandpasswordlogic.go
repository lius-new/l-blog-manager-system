package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/user/internal/svc"
	"github.com/lius-new/blog-backend/api/user/internal/types"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyUserNameAndPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyUserNameAndPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyUserNameAndPasswordLogic {
	return &ModifyUserNameAndPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyUserNameAndPasswordLogic) ModifyUserNameAndPassword(req *types.ModifyUserNameAndPasswordRequest) (resp *types.CreateResponse, err error) {
	_, err = l.svcCtx.Userer.Modify(l.ctx, &user.ModifyUserRequest{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{}, nil
}
