package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
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
func (l *SelectTagByIdLogic) SelectTagById(
	in *content.SelectTagByIdRequest,
) (*content.SelectTagByIdResponse, error) {
	// 找到指定tag
	currentTag, err := l.svcCtx.ModelWithTag.FindOne(l.ctx, in.Id)

	// 判断是否存在错误或者tag是否存在
	if err == rpc.ErrNotFound || currentTag == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &content.SelectTagByIdResponse{
		Id:       currentTag.ID.Hex(),
		Name:     currentTag.Name,
		Articles: currentTag.Articles,
		Visiable: currentTag.Visiable,
	}, nil
}
