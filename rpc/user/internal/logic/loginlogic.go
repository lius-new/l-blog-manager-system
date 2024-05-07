package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/user/internal/svc"
	"github.com/lius-new/blog-backend/rpc/user/user"
	"github.com/lius-new/blog-backend/rpc/utils/utils"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginUserRequest) (*user.LoginResponse, error) {
	// 获取用户的用户信息
	findRes, err := l.svcCtx.Model.FindByUserName(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}
	if !findRes.Status {
		return nil, rpc.ErrInvalidDisabled
	}

	// 对用户表单密码加密后和数据库的密码相比较
	passwordMd5, err := l.svcCtx.Utiler.MD5(l.ctx, &utils.MD5Reqeust{
		Text: in.Password,
	})
	if err != nil {
		return nil, err
	}
	if findRes.Password != passwordMd5.Text {
		return nil, rpc.ErrInvalidPassword
	}

	return &user.LoginResponse{
		Id:       findRes.ID.Hex(),
		Username: findRes.Username,
		Password: findRes.Password,
		Status:   findRes.Status,
		SecretId: findRes.SecretId,
	}, nil
}
