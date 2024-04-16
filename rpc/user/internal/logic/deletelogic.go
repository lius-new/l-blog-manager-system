package logic

import (
	"context"
	"errors"

	"github.com/lius-new/blog-backend/rpc/user/internal/svc"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	count, err := l.svcCtx.Model.Delete(l.ctx, in.Uid)
	if err != nil || count == 0 {
		return nil, errors.New("delete failed")
	}

	return &user.DeleteUserResponse{Uid: in.Uid}, nil
}
