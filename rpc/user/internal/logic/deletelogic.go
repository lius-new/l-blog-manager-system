package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/user/internal/svc"
	"github.com/lius-new/blog-backend/rpc/user/user"
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
	// 删除指定的用户
	count, err := l.svcCtx.Model.Delete(l.ctx, in.Uid)
	// 判断是否删除失败
	if err != nil || count == 0 {
		return nil, rpc.ErrInvalidDeleted
	}
	// 删除生成返回删除的id
	return &user.DeleteUserResponse{Uid: in.Uid}, nil
}
