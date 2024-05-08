package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/user/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/user/model/mongo"
	"github.com/lius-new/blog-backend/rpc/user/user"
	"github.com/lius-new/blog-backend/rpc/utils/utils"

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
	// 检查接口的参数是否存在
	if len(in.Username) == 0 || len(in.Password) == 0 {
		return nil, rpc.ErrRequestParam
	}

	// 查询指定用户名的用户
	findUser, err := l.svcCtx.Model.FindByUserName(l.ctx, in.Username)
	// 判断用户是否存在
	if err != nil && err != rpc.ErrNotFound {
		return nil, err
	}
	// 判断用户是否存在
	if findUser != nil {
		return nil, rpc.ErrInvalidExist
	}

	// utils rpc 服务来加密密码
	passwordMd5, err := l.svcCtx.Utiler.MD5(l.ctx, &utils.MD5Reqeust{
		Text: in.Password,
	})

	// 加密是否存在异常
	if err != nil {
		return nil, err
	}

	// 添加用户
	if err := l.svcCtx.Model.Insert(l.ctx, &model.User{
		Username: in.Username,
		Password: passwordMd5.Text,
		Status:   true,
	}); err != nil {
		return nil, err
	}

	return &user.UserResponse{
		Username: in.Username,
		Password: in.Password,
	}, nil
}
