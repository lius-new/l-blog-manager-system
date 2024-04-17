package logic

import (
	"context"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectTagByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectTagByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectTagByPageLogic {
	return &SelectTagByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据分页获取tag
func (l *SelectTagByPageLogic) SelectTagByPage(in *content.SelectTagByPageRequest) (*content.SelectTagByPageResponse, error) {
	// todo: add your logic here and delete this line

	return &content.SelectTagByPageResponse{}, nil
}
