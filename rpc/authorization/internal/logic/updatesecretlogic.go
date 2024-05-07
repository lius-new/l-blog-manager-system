package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/authorization/model/mongo"
)

type UpdateSecretLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSecretLogic {
	return &UpdateSecretLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新secret
func (l *UpdateSecretLogic) UpdateSecret(
	in *authorization.InsertAndUpdateSecretRequestWithSecret,
) (*authorization.SecretResponseWithSecret, error) {
	if len(in.Id) == 0 {
		return nil, rpc.ErrRequestParam
	}

	// 查询到指定id的secret
	findSecret, err := l.svcCtx.Model.FindOne(l.ctx, in.Id)

	if err != nil && err == rpc.ErrNotFound {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}
  // 检查提交的参数避免提交错误的参数也修改数据库
	if in.SecretInner != findSecret.SecretInner && len(in.SecretInner) != 0 {
		findSecret.SecretInner = in.SecretInner
	}
	if in.SecretOuter != findSecret.SecretOuter && len(in.SecretOuter) != 0 {
		findSecret.SecretOuter = in.SecretOuter
	}
	if in.Uid != findSecret.UserId && len(in.Uid) != 0 {
		findSecret.UserId = in.Uid
	}
	if in.Issuer != findSecret.Issuer && len(in.Issuer) != 0 {
		findSecret.Issuer = in.Issuer
	}
	if in.Expire != findSecret.Expire && in.Expire != 0 {
		findSecret.Expire = in.Expire
	}
  // 更新
	updateRes, err := l.svcCtx.Model.Update(l.ctx, &model.Secret{
		ID:          findSecret.ID,
		SecretInner: findSecret.SecretInner,
		SecretOuter: findSecret.SecretOuter,
		Expire:      findSecret.Expire,
		Issuer:      findSecret.Issuer,
		UserId:      findSecret.UserId,
	})
	if err != nil {
		return nil, err
	} else if updateRes.UpsertedCount == 0 {
		return nil, model.ErrNotFound
	}

	return &authorization.SecretResponseWithSecret{
		Id:          in.Id,
		SecretInner: in.SecretInner,
		SecretOuter: in.SecretOuter,
		Expire:      in.Expire,
		Issuer:      in.Issuer,
		Uid:         in.Uid,
	}, nil
}
