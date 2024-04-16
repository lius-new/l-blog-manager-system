package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/user/internal/svc"
	"github.com/lius-new/blog-backend/api/user/internal/types"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectByPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectByPageLogic {
	return &SelectByPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectByPageLogic) SelectByPage(req *types.SelectPage) (resp *types.SelectPageResponse, err error) {
	selectPageResp, err := l.svcCtx.Userer.SelectByPage(l.ctx, &user.SelectUserByPageRequest{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return &types.SelectPageResponse{
			Status: false,
		}, err
	}
	users := make([]types.UserBackend, 0)
	for _, v := range selectPageResp.Users {
		users = append(users, types.UserBackend{
			Id:       v.Id,
			Username: v.Username,
			Status:   v.Status,
		})

	}

	return &types.SelectPageResponse{
		Data:   users,
		Status: true,
	}, nil
}
