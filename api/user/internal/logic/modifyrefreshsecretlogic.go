package logic

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/lius-new/blog-backend/api"
	"github.com/lius-new/blog-backend/api/user/internal/svc"
	"github.com/lius-new/blog-backend/api/user/internal/types"
	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyRefreshSecretLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyRefreshSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyRefreshSecretLogic {
	return &ModifyRefreshSecretLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyRefreshSecretLogic) ModifyRefreshSecret(req *types.ModifyRefreshSecretRequest) (resp *types.CreateResponse, err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			catchErr := catchErr.(error)
			switch {
			case strings.Contains(catchErr.Error(), api.ErrNotFound.Error()):
				err = errors.New("用户不存在")
			case strings.Contains(catchErr.Error(), api.ErrInvalidObjectId.Error()):
				err = api.ErrInvalidObjectId
			}
		} else if err != nil {
			err = errors.New(strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1))
		}
	}()

	// 查询到指定用户
	userResp, err := l.svcCtx.Userer.Select(l.ctx, &user.SelectUserRequest{
		Uid: req.Id,
	})
	if err != nil {
		panic(err)
	}

	// 生成新的secretID
	generateSecretResp, err := l.svcCtx.Auther.GenerateSecret(
		l.ctx,
		&authorization.GenerateSecretRequestWithSecret{
			Uid:    req.Id,
			Expire: int64(time.Hour * time.Duration(req.Hour)),
			Issuer: "login",
		},
	)
	if err != nil {
		panic(err)
	}

	// 修改当前用户的secretID
	_, err = l.svcCtx.Userer.ModifySecretId(l.ctx, &user.ModifySecretRequest{
		Id:       req.Id,
		SecretId: generateSecretResp.Id,
	})
	if err != nil {
		panic(err)
	}

	// 删除指定用户对应的secretId
	_, err = l.svcCtx.Auther.DeleteSecret(l.ctx, &authorization.DeleteSecretRequestWithSecret{
		Id: userResp.SecretId,
	})

	if err != nil {
		panic(err)
	}

	return &types.CreateResponse{}, nil
}
