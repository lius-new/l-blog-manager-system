package logic

import (
	"context"
	"errors"
	"strings"

	"github.com/lius-new/blog-backend/api"
	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/lius-new/blog-backend/rpc/content/content"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByTagNameWithViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByTagNameWithViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByTagNameWithViewLogic {
	return &GetArticleByTagNameWithViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByTagNameWithViewLogic) GetArticleByTagNameWithView(req *types.GetArticleByTagNameWithViewRequest) (resp *types.GetArticleByTagNameWithViewResponse, err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			var catchErr = catchErr.(error)
			switch {
			case strings.Contains(catchErr.Error(), api.ErrRequestParam.Error()):
				err = errors.New(api.ErrRequestParam.Error())
			case strings.Contains(catchErr.Error(), api.ErrNotFound.Error()):
				err = errors.New(api.ErrInvalidNotFound.Error())
			}
		} else if err != nil {
			err = errors.New(strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1))
		}
	}()
	articleResp, err := l.svcCtx.Content.SelectArtilceByTag(l.ctx, &content.SelectArticleByTagRequest{
		Tag:      req.TagName,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		panic(err)
	}

	// 封装数据
	forLen := len(articleResp.Articles)
	data := make([]types.Article, forLen)
	for i := 0; i < forLen; i++ {
		current := articleResp.Articles[i]
		data[i] = types.Article{
			Id:          current.Id,
			Title:       current.Title,
			Description: current.Desc,
			Tags:        current.Tags,
			UpdateAt:    current.Time,
		}
	}

	return &types.GetArticleByTagNameWithViewResponse{
		Data:  data,
		Total: articleResp.Total,
	}, nil

}
