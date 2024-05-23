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

type GetArticleByIdWithViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByIdWithViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByIdWithViewLogic {
	return &GetArticleByIdWithViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByIdWithViewLogic) GetArticleByIdWithView(req *types.GetArticleByIdWithViewRequest) (resp *types.GetArticleByIdWithViewResponse, err error) {
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
	// 查询数据
	articleResp, err := l.svcCtx.Content.SelectArtilceById(l.ctx, &content.SelectArticleByIdRequest{
		Id: req.Id,
	})
	if err != nil {
		panic(err)
	}

	return &types.GetArticleByIdWithViewResponse{
		Id:          articleResp.Id,
		Title:       articleResp.Title,
		Description: articleResp.Desc,
		Content:     articleResp.Content,
		Tags:        articleResp.Tags,
		Covers:      articleResp.Covers,
		UpdateAt:    articleResp.Time,
	}, nil
}
