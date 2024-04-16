package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/user/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/user/model/mongo"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLogic {
	return &InsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Insert: 添加用户
func (l *InsertLogic) Insert(in *user.InsertUserRequest) (*user.UserResponse, error) {
	if len(in.Username) == 0 || len(in.Password) == 0 {
		return nil, rpc.ErrRequestParam
	}

	// 判断用户是否存在
	findUser, err := l.svcCtx.Model.FindByUserName(l.ctx, in.Username)
	if err != nil && err != rpc.ErrNotFound {
		return nil, err
	}
	if findUser != nil {
		return nil, rpc.ErrInvalidExist
	}

	if err := l.svcCtx.Model.Insert(l.ctx, &model.User{
		Username: in.Username,
		Password: in.Password,
		Status:   true,
	}); err != nil {
		return nil, err
	}

	return &user.UserResponse{
		Username: in.Username,
		Password: in.Password,
	}, nil
}
