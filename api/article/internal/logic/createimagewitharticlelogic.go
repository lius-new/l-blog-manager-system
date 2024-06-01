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

type CreateImageWithArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateImageWithArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateImageWithArticleLogic {
	return &CreateImageWithArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 添加文章中的图片，有些文章内部会包含图片. TODO: 这个接口会将数据放到covers数据集合中(mongodb), 事实上cover并不是存放这些图片的地方现在是图省事。
func (l *CreateImageWithArticleLogic) CreateImageWithArticle(req *types.CreateImageWithArticleRequest) (resp *types.CreateImageWithArticleResponse, err error) {

	// defer recover错误信息
	defer func() {
		if catchErr := recover(); catchErr != nil {
			var catchErr = catchErr.(error)
			switch {
			case strings.Contains(catchErr.Error(), api.ErrRequestParam.Error()):
				err = errors.New(api.ErrRequestParam.Error())
			case strings.Contains(catchErr.Error(), api.ErrInvalidExist.Error()):
				err = errors.New(api.ErrInvalidExist.Error())
			}
		} else if err != nil {
			err = errors.New(strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1))
		}
	}()

	createResp, err := l.svcCtx.Content.CreateCovers(l.ctx, &content.CreateCoversRequest{
		Content: req.Contents,
	})
	if err != nil {
		panic(err)
	}

	return &types.CreateImageWithArticleResponse{
		Hashs: createResp.Hashs,
	}, nil
}
