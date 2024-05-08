package logic

import (
	"context"
	"errors"
	"strings"

	"github.com/lius-new/blog-backend/api"
	"github.com/lius-new/blog-backend/api/user/internal/svc"
	"github.com/lius-new/blog-backend/api/user/internal/types"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			var catchErr = catchErr.(error)
			switch {
			case strings.Contains(catchErr.Error(), api.ErrInvalidExist.Error()):
				err = errors.New("用户名已存在")
			}
		} else if err != nil {
			err = errors.New(strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1))
		}
	}()

	// 查询用户
	uResp, err := l.svcCtx.Userer.SelectByPage(l.ctx, &user.SelectUserByPageRequest{
		PageNum:  0,
		PageSize: 10,
	})
	if err != nil {
		panic(err)
	}
	if uResp.Total > 1 {
		return nil, errors.New("请联系管理员")
	}
	newUser := &user.InsertUserRequest{
		Username: req.Username,
		Password: req.Password,
	}

	_, err = l.svcCtx.Userer.Insert(l.ctx, newUser)

	if err != nil {
		panic(err)
	}

	return &types.CreateResponse{}, nil
}
