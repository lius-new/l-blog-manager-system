package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/cover"
	"github.com/lius-new/blog-backend/rpc/utils/utiler"
)

type CreateCoverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCoverLogic {
	return &CreateCoverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// **cover **
func (l *CreateCoverLogic) CreateCover(
	in *content.CreateCoverRequest,
) (*content.CreateCoverResponse, error) {
	if len(in.Content) == 0 {
		return nil, rpc.ErrRequestParam
	}

	hashResp, err := l.svcCtx.Utiler.HashWithString(l.ctx, &utiler.HashWithStringReqeust{
		Content: in.Content,
	})
	if err != nil {
		return nil, err
	}

	// 判断指定hash的图片是否存在
	c, err := l.svcCtx.ModelWithCover.FindOneByHash(l.ctx, hashResp.Content)
	// 存在就返回id
	if c != nil {
		return &content.CreateCoverResponse{
			Id:   c.ID.Hex(),
			Hash: c.Hash,
		}, nil
	}
	// 如果存在错误且错误不是NotFound那么就抛出
	if err != nil && err != rpc.ErrNotFound {
		return nil, err
	}

	// 保存base64图片到数据库中。
	id, err := l.svcCtx.ModelWithCover.InsertReturnId(l.ctx, &model.Cover{
		Content: in.Content,
		Hash:    hashResp.Content,
	})
	if err != nil {
		return nil, err
	}
	// 保存到数据库后再返回id
	return &content.CreateCoverResponse{
		Id:   id,
		Hash: hashResp.Content,
	}, nil
}
