package logic

import (
	"context"

	"github.com/lius-new/blog-backend/api/user/internal/svc"
	"github.com/lius-new/blog-backend/api/user/internal/types"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	newUser := &user.InsertUserRequest{
		Username: req.Username,
		Password: req.Password,
	}

	insertResponse, err := l.svcCtx.Userer.Insert(l.ctx, newUser)

	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Data: types.UserBackend{
			Id:       insertResponse.Id,
			Username: insertResponse.Username,
			Status:   insertResponse.Status,
		},
	}, nil
}
