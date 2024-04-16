package logic

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/authorization/model/mongo"
	"github.com/lius-new/blog-backend/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateSecretLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateSecretLogic {
	return &GenerateSecretLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func generateRandomKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

// 为指定用户生成secret
func (l *GenerateSecretLogic) GenerateSecret(in *authorization.GenerateSecretRequestWithSecret) (*authorization.SecretResponseWithSecret, error) {
	if len(in.Uid) == 0 || len(in.Issuer) == 0 || in.Expire <= 0 {
		return nil, rpc.ErrRequestParam
	}

	// 检查要求生成secret的用户是否存在
	_, err := l.svcCtx.Userer.Select(l.ctx, &user.SelectUserRequest{
		Uid: in.Uid,
	})

	if err != nil && err == rpc.ErrNotFound {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// 检查用户是否已经存在了secret
	secret, err := l.svcCtx.Model.FindByUID(l.ctx, in.Uid)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	// 如果已经存在了, 那么就删除原本的凭证, 保留现在的.
	if secret != nil {
		l.svcCtx.Model.DeleteByUID(l.ctx, in.Uid)
	}

	secretInner, _ := generateRandomKey(32)
	secretOuter, _ := generateRandomKey(32)
	secret = &model.Secret{
		SecretInner: secretInner,
		SecretOuter: secretOuter,
		Expire:      time.Now().UnixNano() + in.Expire,
		Issuer:      in.Issuer,
		UserId:      in.Uid,
	}
	objectId, err := l.svcCtx.Model.Insert(l.ctx, secret)
	if err != nil {
		return nil, err
	}

	return &authorization.SecretResponseWithSecret{
		Id:          objectId.Hex(),
		SecretInner: secretInner,
		SecretOuter: secretOuter,
		Expire:      in.Expire,
		Issuer:      in.Issuer,
		Uid:         in.Uid,
	}, nil
}