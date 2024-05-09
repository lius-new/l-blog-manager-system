package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/user/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/user/model/mongo"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyUserNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyUserNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyUserNameLogic {
	return &ModifyUserNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyUserNameLogic) ModifyUserName(in *user.ModifyUserNameRequest) (*user.ModifyUserNameRequest, error) {
	if len(in.Id) == 0 || len(in.Username) == 0 {
		return nil, rpc.ErrRequestParam
	}
	findUser, err := l.svcCtx.Model.FindOne(l.ctx, in.Id)
	// 判断用户是否存在
	if err == rpc.ErrNotFound {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	resp, err := l.svcCtx.Model.Update(l.ctx, &model.User{
		ID:       findUser.ID,
		Username: in.Username,
	})
	if err != nil {
		return nil, err
	}
	if resp.ModifiedCount == 0 {
		return nil, rpc.ErrNotFound
	}
	return &user.ModifyUserNameRequest{}, nil
}
