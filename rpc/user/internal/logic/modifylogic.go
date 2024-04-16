package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/user/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/user/model/mongo"
	"github.com/lius-new/blog-backend/rpc/user/user"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zeromicro/go-zero/core/logx"
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
	id, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}

	resp, err := l.svcCtx.Model.Update(l.ctx, &model.User{
		ID:       id,
		Username: in.Username,
		Password: in.Password,
		Status:   in.Status,
		SecretId: in.SecretId,
	})
	if err != nil {
		return nil, err
	}
	if resp.ModifiedCount == 0 {
		return nil, rpc.ErrNotFound
	}

	return &user.ModifyUserResponse{
		Username: in.Username,
		Password: in.Password,
		Status:   in.Status,
		SecretId: in.SecretId,
	}, nil

}
