package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/lius-new/blog-backend/rpc/utils/internal/svc"
	"github.com/lius-new/blog-backend/rpc/utils/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type MD5Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMD5Logic(ctx context.Context, svcCtx *svc.ServiceContext) *MD5Logic {
	return &MD5Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MD5Logic) MD5(in *utils.MD5Reqeust) (*utils.MD5Reponse, error) {
	md5er := md5.New()
	md5er.Write([]byte(in.Text))
	md5Res := md5er.Sum(nil)

	return &utils.MD5Reponse{
		Text: hex.EncodeToString(md5Res),
	}, nil
}
