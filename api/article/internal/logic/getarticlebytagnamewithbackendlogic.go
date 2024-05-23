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

type GetArticleByTagNameWithBackendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByTagNameWithBackendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByTagNameWithBackendLogic {
	return &GetArticleByTagNameWithBackendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByTagNameWithBackendLogic) GetArticleByTagNameWithBackend(req *types.GetArticleByTagNameWithBackendRequest) (resp *types.GetArticleByTagNameWithBackendResponse, err error) {
	// todo: add your logic here and delete this line
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
	data := make([]types.Data, forLen)
	for i := 0; i < forLen; i++ {
		current := articleResp.Articles[i]
		data[i] = types.Data{
			Id:          current.Id,
			Title:       current.Title,
			Description: current.Desc,
			Tags:        current.Tags,
			Covers:      current.Covers,
			UpdateAt:    current.Time,
		}
	}

	return &types.GetArticleByTagNameWithBackendResponse{
		Data:  data,
		Total: articleResp.Total,
	}, nil
}
