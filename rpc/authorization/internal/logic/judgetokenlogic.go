package logic

import (
	"context"
	"time"

	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/jwt"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJudgeTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeTokenLogic {
	return &JudgeTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 校验token
func (l *JudgeTokenLogic) JudgeToken(in *authorization.JudgeJwtRequestWithJwt) (*authorization.JudgeJwtResponseWithJwt, error) {
	secret, err := l.svcCtx.Model.FindOne(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}
	expire := time.Now().Add(time.Duration(secret.Expire))
	jwtUtil := jwt.NewJwtUtil(secret.SecretInner, secret.SecretOuter, expire, secret.Issuer)
	cliasm, err := jwtUtil.ParseJwtToken(in.Token)

	if err != nil || len(cliasm.ID) == 0 {
		return nil, err
	}

	return &authorization.JudgeJwtResponseWithJwt{
		Id: cliasm.ID,
	}, nil
}
