package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/user/internal/svc"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectByNameLogic {
	return &SelectByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectByNameLogic) SelectByName(in *user.SelectUserByUsernameRequest) (*user.UserResponse, error) {
	findUser, err := l.svcCtx.Model.FindByUserName(l.ctx, in.Username)

	if err == rpc.ErrNotFound {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &user.UserResponse{
		Id:       findUser.ID.Hex(),
		Username: findUser.Username,
		Status:   findUser.Status,
		SecretId: findUser.SecretId,
	}, nil
}
