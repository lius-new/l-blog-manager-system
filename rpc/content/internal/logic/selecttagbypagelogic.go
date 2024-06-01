package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
)

type SelectTagByPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectTagByPageLogic(
	ctx context.Context,
	svcCtx *svc.ServiceContext,
) *SelectTagByPageLogic {
	return &SelectTagByPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据分页获取tag
func (l *SelectTagByPageLogic) SelectTagByPage(
	in *content.SelectTagByPageRequest,
) (*content.SelectTagByPageResponse, error) {
	// 判断分页要求的参数
	if in.PageSize == 0 || in.PageNum == 0 {
		return nil, rpc.ErrRequestParam
	}

	tags, total, err := l.svcCtx.ModelWithTag.FindByPage(
		l.ctx,
		in.PageNum,
		in.PageSize,
		in.HideShow,
	)
	if err != nil {
		return nil, err
	}

	forLen := len(tags)
	resptags := make([]*content.SelectTag, forLen)

	for i := 0; i < forLen; i++ {
		currentTag := tags[i]
		resptags[i] = &content.SelectTag{
			Id:   currentTag.ID.Hex(),
			Name: currentTag.Name,
		}
	}

	return &content.SelectTagByPageResponse{
		Tags:  resptags,
		Total: total,
	}, nil
}
