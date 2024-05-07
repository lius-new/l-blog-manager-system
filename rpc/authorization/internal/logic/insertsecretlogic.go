package logic

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/authorization/model/mongo"
)

type InsertSecretLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertSecretLogic {
	return &InsertSecretLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增secret
func (l *InsertSecretLogic) InsertSecret(
	in *authorization.InsertAndUpdateSecretRequestWithSecret,
) (*authorization.SecretResponseWithSecret, error) {
	// 查询是否存在
	secret, err := l.svcCtx.Model.FindByUID(l.ctx, in.Uid)
	if err != nil {
		return nil, err
	}
	// 如果已经存在就抛出
	if secret != nil {
		return nil, model.ErrExist
	}

	secret = &model.Secret{
		SecretInner: in.SecretInner,
		SecretOuter: in.SecretOuter,
		Expire:      in.Expire,
		Issuer:      in.Issuer,
		UserId:      in.Uid,
	}

	objectId, err := l.svcCtx.Model.Insert(l.ctx, secret)
	if err != nil {
		return nil, err
	}

	return &authorization.SecretResponseWithSecret{
		Id:          objectId.String(),
		SecretInner: in.SecretInner,
		SecretOuter: in.SecretOuter,
		Expire:      in.Expire,
		Issuer:      in.Issuer,
		Uid:         in.Uid,
	}, nil
}
