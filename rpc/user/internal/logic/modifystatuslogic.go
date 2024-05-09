package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/user/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/user/model/mongo"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyStatusLogic {
	return &ModifyStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyStatusLogic) ModifyStatus(in *user.ModifyUserStatusRequest) (*user.ModifyUserStatusRequest, error) {
	if len(in.Id) == 0 {
		return nil, rpc.ErrRequestParam
	}
	findUser, err := l.svcCtx.Model.FindOne(l.ctx, in.Id)
	// 判断用户是否存在
	if err == rpc.ErrNotFound {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	// 因为我直接调用Update方法总是无法修改status, 所以写了个类似的方法。
	resp, err := l.svcCtx.Model.UpdateStatus(l.ctx, &model.User{
		ID:     findUser.ID,
		Status: in.Status,
	})
	if err != nil {
		return nil, err
	}
	if resp.ModifiedCount == 0 {
		return nil, rpc.ErrNotFound
	}

	return &user.ModifyUserStatusRequest{}, nil
}
