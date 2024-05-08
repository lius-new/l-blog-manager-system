package logic

import (
	"context"
	"time"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/jwt"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/svc"
	"github.com/lius-new/blog-backend/rpc/user/user"

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
	// 外部secret放在svcCtx上，实际上是查询redis中的secret.
	secretOuter := l.svcCtx.GetSecretOuter()
	// 解析外层
	outerClaims, err := jwt.NewJwtUtil("", secretOuter, time.Now(), "").ParseJwtTokenOuter(in.Token)
	if err != nil {
		return nil, err
	}
	// 判断解析外层token内容key是否为空
	if len(outerClaims.Key) == 0 || len(outerClaims.Value) == 0 {
		return nil, err
	}

	// 查询到指定用户
	findUser, err := l.svcCtx.Userer.SelectByName(l.ctx, &user.SelectUserByUsernameRequest{Username: outerClaims.Key})
	if err != nil {
		return nil, err
	}
	// 查询指定secret并判断是否存在
	innerSecret, err := l.svcCtx.Model.FindOne(l.ctx, findUser.SecretId)
	if err == rpc.ErrNotFound {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// TODO: expire 在解析token的情况下也许不需要设置，如果要设置那么这里的设置也是错误的,应该是生成时候的时间加上expire
	expire := time.Now().Add(time.Duration(innerSecret.Expire))
	// 检验
	cliasm, err := jwt.NewJwtUtil(innerSecret.Secret, "", expire, innerSecret.Issuer).ParseJwtTokenInner(outerClaims.Value)

	if len(cliasm.Key) == 0 || len(cliasm.Value) == 0 {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &authorization.JudgeJwtResponseWithJwt{
		Id: cliasm.Key,
	}, nil
}
