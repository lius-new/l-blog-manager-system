package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectSecretLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectSecretLogic {
	return &SelectSecretLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询secret
func (l *SelectSecretLogic) SelectSecret(in *authorization.SelectSecretRequestWithSecret) (*authorization.SecretResponseWithSecret, error) {
	secret, err := l.svcCtx.Model.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &authorization.SecretResponseWithSecret{
		Id:          secret.ID.String(),
		SecretInner: secret.SecretInner,
		SecretOuter: secret.SecretOuter,
		Expire:      secret.Expire,
		Issuer:      secret.Issuer,
		Uid:         secret.UserId,
	}, nil
}
