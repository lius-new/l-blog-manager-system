package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/user/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/user/model/mongo"
	"github.com/lius-new/blog-backend/rpc/user/user"
	"github.com/lius-new/blog-backend/rpc/utils/utils"
)

type ModifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyLogic {
	return &ModifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyLogic) Modify(in *user.ModifyUserRequest) (*user.ModifyUserResponse, error) {
	// 如果用户id不存在就返回nil,err
	if len(in.Id) == 0 {
		return nil, rpc.ErrRequestParam
	}

	// 查询到，然后判断是否更新，如果跟新就更新数据库，否则不更新数据库
	findUser, err := l.svcCtx.Model.FindOne(l.ctx, in.Id)

	// 判断用户是否存在
	if err == rpc.ErrNotFound {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// 比较用户名是否被更新
	if findUser.Username != in.Username && len(in.Username) != 0 {
		findUser.Username = in.Username
	}

	// 比较密码是否更新
	if len(in.Password) != 0 {
		passowordMD5, err := l.svcCtx.Utiler.MD5(l.ctx, &utils.MD5Reqeust{Text: in.Password})
		if err != nil {
			return nil, err
		}
		if findUser.Password != passowordMD5.Text {
			findUser.Password = passowordMD5.Text
		}
	}

	// 比较密钥是否更新
	if findUser.SecretId != in.SecretId && len(in.SecretId) != 0 {
		findUser.SecretId = in.SecretId
	}

	// 更新用户信息
	resp, err := l.svcCtx.Model.Update(l.ctx, &model.User{
		ID:       findUser.ID,
		Username: findUser.Username,
		Status:   in.Status,
		SecretId: findUser.SecretId,
	})
	// 更新操作是否存在err
	if err != nil {
		return nil, err
	}
	if resp.ModifiedCount == 0 {
		return nil, rpc.ErrNotFound
	}

	return &user.ModifyUserResponse{
		Username: findUser.Username,
		Status:   in.Status,
		SecretId: findUser.SecretId,
	}, nil
}
