package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/authorization/model/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zeromicro/go-zero/core/logx"
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
func (l *UpdateSecretLogic) UpdateSecret(in *authorization.InsertAndUpdateSecretRequestWithSecret) (*authorization.SecretResponseWithSecret, error) {
	id, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, model.ErrInvalidObjectId
	}

	updateRes, err := l.svcCtx.Model.Update(l.ctx, &model.Secret{
		ID:          id,
		SecretInner: in.SecretInner,
		SecretOuter: in.SecretOuter,
		Expire:      in.Expire,
		Issuer:      in.Issuer,
		UserId:      in.Uid,
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
