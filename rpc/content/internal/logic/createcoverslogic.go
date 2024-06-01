package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type CreateCoversLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCoversLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCoversLogic {
	return &CreateCoversLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建多个Cover
func (l *CreateCoversLogic) CreateCovers(
	in *content.CreateCoversRequest,
) (*content.CreateCoversResponse, error) {
	createCoverLogic := NewCreateCoverLogic(l.ctx, l.svcCtx)

	forLen := len(in.Content)
	ids := make([]string, forLen)
	hashs := make([]string, forLen)

	for i := 0; i < forLen; i++ {
		createResp, err := createCoverLogic.CreateCover(&content.CreateCoverRequest{
			Content: in.Content[i],
		})
		// TODO: 参考创建的逻辑好像一般不会出错, 如果出错后面再说吧^-^
		if err != nil {
			return nil, err
		}

		ids[i] = createResp.GetId()
		hashs[i] = createResp.Hash
	}

	return &content.CreateCoversResponse{
		Ids:   ids,
		Hashs: hashs,
	}, nil
}
