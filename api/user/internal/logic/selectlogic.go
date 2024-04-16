package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/user/internal/svc"
	"github.com/lius-new/blog-backend/api/user/internal/types"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectLogic {
	return &SelectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectLogic) Select(req *types.IdRequest) (resp *types.CreateResponse, err error) {
	selectResp, err := l.svcCtx.Userer.Select(l.ctx, &user.SelectUserRequest{
		Uid: req.Id,
	})
	if err != nil {
		return &types.CreateResponse{
			Status: false,
		}, err
	}

	return &types.CreateResponse{
		Data: types.UserBackend{
			Id:       req.Id,
			Username: selectResp.Username,
			Status:   selectResp.Status,
		},
		Status: true,
	}, nil
}
