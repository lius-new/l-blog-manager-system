package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type CreateTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTagsLogic {
	return &CreateTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建标签，如果标签存在也会查询出来放入到返回结果，如果不存在就创建再放入到返回结果。
func (l *CreateTagsLogic) CreateTags(
	in *content.CreateTagsRequest,
) (*content.CreateTagsResponse, error) {
	createTagLogic := NewCreateTagLogic(l.ctx, l.svcCtx)

	ids := make([]string, len(in.Names))
	for _, v := range in.Names {
		createResp, err := createTagLogic.CreateTag(&content.CreateTagRequest{
			Name: v,
		})
		// TODO: 参考创建的逻辑好像一般不会出错, 如果出错后面再说吧^-^
		if err != nil {
			return nil, err
		}

		ids = append(ids, createResp.GetId())
	}

	return &content.CreateTagsResponse{
		Ids: ids,
	}, nil
}
