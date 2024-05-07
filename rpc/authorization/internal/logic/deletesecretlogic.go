package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
	"github.com/lius-new/blog-backend/rpc/authorization/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/authorization/model/mongo"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSecretLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSecretLogic {
	return &DeleteSecretLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除secret
func (l *DeleteSecretLogic) DeleteSecret(in *authorization.DeleteSecretRequestWithSecret) (*authorization.DeleteSecretResponseWithSecret, error) {
  // 删除secret
	deleteCount, err := l.svcCtx.Model.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	} else if deleteCount == 0 {
		return nil, model.ErrNotFound
	}

	return &authorization.DeleteSecretResponseWithSecret{
		Id: in.Id,
	}, nil
}
