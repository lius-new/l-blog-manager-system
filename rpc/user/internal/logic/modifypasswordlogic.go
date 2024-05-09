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

type ModifyPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyPasswordLogic {
	return &ModifyPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyPasswordLogic) ModifyPassword(in *user.ModifyPasswordRequest) (*user.ModifyPasswordResponse, error) {
	if len(in.Id) == 0 || len(in.Password) == 0 {
		return nil, rpc.ErrRequestParam
	}

	findUser, err := l.svcCtx.Model.FindOne(l.ctx, in.Id)
	// 判断用户是否存在
	if err == rpc.ErrNotFound {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
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

	resp, err := l.svcCtx.Model.Update(l.ctx, &model.User{
		ID:       findUser.ID,
		Password: findUser.Password,
	})
	if err != nil {
		return nil, err
	}
	if resp.ModifiedCount == 0 {
		return nil, rpc.ErrNotFound
	}

	return &user.ModifyPasswordResponse{}, nil
}
