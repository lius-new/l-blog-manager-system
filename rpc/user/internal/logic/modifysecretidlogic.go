package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/user/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/user/model/mongo"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifySecretIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifySecretIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifySecretIdLogic {
	return &ModifySecretIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifySecretIdLogic) ModifySecretId(in *user.ModifySecretRequest) (*user.ModifySecretRequest, error) {
	if len(in.Id) == 0 || len(in.SecretId) == 0 {
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
		SecretId: in.SecretId,
	})

	if err != nil {
		return nil, err
	}
	if resp.ModifiedCount == 0 {
		return nil, rpc.ErrNotFound
	}
	return &user.ModifySecretRequest{}, nil
}
