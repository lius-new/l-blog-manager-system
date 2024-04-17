package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectTagByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectTagByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectTagByIdLogic {
	return &SelectTagByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据id获取tag
func (l *SelectTagByIdLogic) SelectTagById(in *content.SelectTagByIdRequest) (*content.SelectTagByIdResponse, error) {
	tag, err := l.svcCtx.ModelWithTag.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &content.SelectTagByIdResponse{
		Id:       tag.ID.Hex(),
		Name:     tag.Name,
		Articles: tag.Articles,
		Visiable: tag.Visiable,
	}, nil
}
