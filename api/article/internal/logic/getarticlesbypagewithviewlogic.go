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

type GetArticlesByPageWithViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticlesByPageWithViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlesByPageWithViewLogic {
	return &GetArticlesByPageWithViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticlesByPageWithViewLogic) GetArticlesByPageWithView(req *types.GetArticleByPageWithViewRequest) (resp *types.GetArticleByPageWithViewResponse, err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			var catchErr = catchErr.(error)
			switch {
			case strings.Contains(catchErr.Error(), api.ErrRequestParam.Error()):
				err = errors.New(api.ErrRequestParam.Error())
			}
		} else if err != nil {
			err = errors.New(strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1))
		}
	}()

	// 调用rpc获取结果
	articlesResp, err := l.svcCtx.Content.SelectArtilceByPage(l.ctx, &content.SelectArticleByPageRequest{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		HideShow: false,
	})
	if err != nil {
		panic(err)
	}

	// 封装数据
	forLen := len(articlesResp.Articles)
	data := make([]types.Article, forLen)
	for i := 0; i < forLen; i++ {
		current := articlesResp.Articles[i]
		data[i] = types.Article{
			Id:          current.Id,
			Title:       current.Title,
			Description: current.Desc,
			Tags:        current.Tags,
			UpdateAt:    current.Time,
		}
	}

	return &types.GetArticleByPageWithViewResponse{
		Data:  data,
		Total: articlesResp.Total,
	}, nil

}
