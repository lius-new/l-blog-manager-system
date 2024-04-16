package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/user/internal/svc"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectLogic {
	return &SelectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectLogic) Select(in *user.SelectUserRequest) (*user.UserResponse, error) {
	if len(in.Uid) == 0 {
		return nil, rpc.ErrRequestParam
	}

	selectUser, err := l.svcCtx.Model.FindOne(l.ctx, in.Uid)

	if err == rpc.ErrNotFound {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &user.UserResponse{
		Username: selectUser.Username,
		Password: selectUser.Password,
		SecretId: selectUser.SecretId,
	}, nil
}
