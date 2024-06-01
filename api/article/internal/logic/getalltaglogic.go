package logic

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/lius-new/blog-backend/api"
	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/rpc/content/content"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllTagLogic {
	return &GetAllTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllTagLogic) GetAllTag(req *types.GetAllTagRequest) (resp *types.GetAllTagResponse, err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			var catchErr = catchErr.(error)
			switch {
			case strings.Contains(catchErr.Error(), api.ErrNotFound.Error()):
				err = errors.New(api.ErrInvalidNotFound.Error())
			}
		} else if err != nil {
			err = errors.New(strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1))
		}
	}()

	tagResp, err := l.svcCtx.Content.SelectTagByPage(l.ctx, &content.SelectTagByPageRequest{
		PageNum:  1,
		PageSize: -1,
		HideShow: true,
	})

	if err != nil {
		panic(err)
	}

	forLen := len(tagResp.Tags)

	respData := make([]types.TagResponse, forLen)

	for i := 0; i < forLen; i++ {
		currentTag := tagResp.Tags[i]
		fmt.Println(currentTag)
		respData[i] = types.TagResponse{
			Id:   currentTag.Id,
			Name: currentTag.Name,
		}
	}

	return &types.GetAllTagResponse{
		Data: respData,
	}, nil
}
