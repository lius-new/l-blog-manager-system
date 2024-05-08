package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/user/internal/svc"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectByPageLogic {
	return &SelectByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SelectByPageLogic) SelectByPage(in *user.SelectUserByPageRequest) (*user.UsersResponse, error) {
	users, total, err := l.svcCtx.Model.FindByPage(l.ctx, in.PageNum, in.PageSize)
	if err != nil {
		return nil, err
	}

	res := make([]*user.UserResponse, 0)

	for _, v := range users {
		res = append(res, &user.UserResponse{
			Id:       v.ID.Hex(),
			Username: v.Username,
			Password: v.Password,
			Status:   v.Status,
			SecretId: v.SecretId,
		})
	}

	return &user.UsersResponse{Users: res, Total: total}, nil
}
