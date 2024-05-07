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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 统一错误处理
	defer func() {
		if catchErr := recover(); catchErr != nil {
			var catchErr = catchErr.(error)
			switch {
			case strings.Contains(catchErr.Error(), api.ErrNotFound.Error()):
				err = errors.New("用户不存在")
			case strings.Contains(catchErr.Error(), api.ErrInvalidDisabled.Error()):
				err = errors.New("用户被禁用")
			case strings.Contains(catchErr.Error(), api.ErrInvalidObjectId.Error()):
				err = api.ErrInvalidObjectId
			case strings.Contains(catchErr.Error(), "password error"):
				err = errors.New("用户密码错误")
			}
		}
		if err != nil {
			err = errors.New(strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1))
		}
	}()

	// 调用登陆rpc接口, 登陆成功返回用户需要的数据.
	loginResp, err := l.svcCtx.Userer.Login(l.ctx, &user.LoginUserRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		panic(err)
	}

	// 检查是否存在, 也就是说, 也许secretid有, 但是secret数据却不存在的情况.
	_, err = l.svcCtx.Auther.SelectSecret(l.ctx, &authorization.SelectSecretRequestWithSecret{Id: loginResp.SecretId})

	// 登陆成功后判断用户是否包含凭证id, 如果没有就要生成并修改用户数据, 为其添加凭证id. 凭证id可查询用户token加密的密钥
	if len(loginResp.SecretId) == 0 || (err != nil && strings.Contains(err.Error(), api.ErrNotFound.Error())) {
		generateSecretResp, err := l.svcCtx.Auther.GenerateSecret(l.ctx, &authorization.GenerateSecretRequestWithSecret{
			Uid:    loginResp.Id,
			Expire: int64(time.Hour * 6),
			Issuer: "login",
		})
		if err != nil {
			panic(err)
		}

    // 修改当前用户的凭证Id
		_, err = l.svcCtx.Userer.Modify(l.ctx, &user.ModifyUserRequest{
			Id:       loginResp.Id,
			Username: loginResp.Username,
			Password: loginResp.Password,
			Status:   loginResp.Status,
			SecretId: generateSecretResp.Id,
		})

		if err != nil {
			panic(err)
		}
		loginResp.SecretId = generateSecretResp.GetId()
	}

  // 生成token
	generateResp, err := l.svcCtx.Auther.GenerateToken(l.ctx, &authorization.GenerateJwtRequestWithJwt{
		Id:       loginResp.SecretId,
		Uid:      loginResp.Id,
		Uesrname: loginResp.Username,
	})
	if err != nil {
		panic(err)
	}

	return &types.LoginResponse{
		Token: generateResp.Token,
	}, nil
}
