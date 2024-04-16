package logic

import (
	"context"
	"time"

	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/jwt"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生成token
func (l *GenerateTokenLogic) GenerateToken(in *authorization.GenerateJwtRequestWithJwt) (*authorization.GenerateJwtResponseWithJwt, error) {
	secret, err := l.svcCtx.Model.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	expire := time.Now().Add(time.Duration(secret.Expire))
	jwtUtil := jwt.NewJwtUtil(secret.SecretInner, secret.SecretOuter, expire, secret.Issuer)
	token, err := jwtUtil.GenerateJwtToken(in.Uid, in.Uesrname)
	if err != nil {
		return nil, err
	}
	return &authorization.GenerateJwtResponseWithJwt{
		Token: token,
	}, nil
}
