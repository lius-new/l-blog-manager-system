package logic

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/jwt"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/svc"
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
func (l *GenerateTokenLogic) GenerateToken(
	in *authorization.GenerateJwtRequestWithJwt,
) (*authorization.GenerateJwtResponseWithJwt, error) {
	// 判断请求参数是否异常
	if len(in.Id) == 0 || len(in.Uesrname) == 0 || len(in.Uid) == 0 {
		return nil, rpc.ErrRequestParam
	}

	// 查询到用户对应的secretID
	secret, err := l.svcCtx.Model.FindOne(l.ctx, in.Id)

	// 判断是否存在
	if err == rpc.ErrNotFound {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// 超时的时间
	expire := time.Now().Add(time.Duration(secret.Expire))

	// 从redis中获取到的secret
	secretOuter := l.svcCtx.GetSecretOuter()
	// jwt util struct
	jwtUtil := jwt.NewJwtUtil(secret.Secret, secretOuter, expire, secret.Issuer)
	// 生成token
	token, err := jwtUtil.GenerateJwtToken(in.Uid, in.Uesrname)

	// 是否存在生成的错误信息
	if err != nil {
		return nil, err
	}
	return &authorization.GenerateJwtResponseWithJwt{
		Token: token,
	}, nil
}
