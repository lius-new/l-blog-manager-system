package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type DeleteTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagLogic {
	return &DeleteTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除tag
func (l *DeleteTagLogic) DeleteTag(
	in *content.DeleteTagRequest,
) (*content.DeleteTagResponse, error) {
	count, err := l.svcCtx.ModelWithTag.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &content.DeleteTagResponse{Count: count}, nil
}
