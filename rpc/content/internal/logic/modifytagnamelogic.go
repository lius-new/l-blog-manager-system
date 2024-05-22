package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/tag"
)

type ModifyTagNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyTagNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyTagNameLogic {
	return &ModifyTagNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改tag name
func (l *ModifyTagNameLogic) ModifyTagName(
	in *content.ModifyTagNameRequest,
) (*content.ModifyTagNameResponse, error) {
	// 判断指定tag是否存在
	currentTag, err := l.svcCtx.ModelWithTag.FindOne(l.ctx, in.Id)

	// 判断是否存在错误或者tag是否存在
	if err == rpc.ErrNotFound || currentTag == nil {
		return nil, rpc.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// TODO: 搞不懂为什么结果体只设置指定属性那么其他属性就会设置为对应零值，日志也显示只修改了指定属性而没有修改其他属性呀
	// 更新
	_, err = l.svcCtx.ModelWithTag.Update(l.ctx, &model.Tag{
		ID:       currentTag.ID,
		Name:     in.Name,
		Articles: currentTag.Articles,
		Visiable: currentTag.Visiable,
	})
	if err != nil {
		return nil, err
	}

	return &content.ModifyTagNameResponse{}, nil
}
