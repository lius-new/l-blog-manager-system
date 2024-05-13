package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lius-new/blog-backend/rpc"
	"github.com/lius-new/blog-backend/rpc/content/content"
	"github.com/lius-new/blog-backend/rpc/content/internal/svc"
	model "github.com/lius-new/blog-backend/rpc/content/model/mongo/article"
)

type CreateArtilceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateArtilceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArtilceLogic {
	return &CreateArtilceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// * article create *
func (l *CreateArtilceLogic) CreateArtilce(
	in *content.CreateArticleRequest,
) (*content.CreateArticleResponse, error) {
	// 判断几个关键参数是否为空
	if len(in.Title) == 0 || len(in.Desc) == 0 || len(in.Content) == 0 || len(in.Tags) == 0 {
		return nil, rpc.ErrRequestParam
	}

	// 搜索要创建文章的文章标题
	a, err := l.svcCtx.ModelWithArticle.FindByTitle(l.ctx, in.Title)
	// 如果a不等nil就意味文章已经存在
	if a != nil {
		return nil, rpc.ErrInvalidExist
	}
	// 如果存在错误且错误不是NotFound那么就抛出
	if err != nil && err != rpc.ErrNotFound {
		return nil, err
	}

	// 创建标签
	tagIds, err := NewCreateTagsLogic(l.ctx, l.svcCtx).CreateTags(&content.CreateTagsRequest{
		Names: in.Tags,
	})
	if err != nil {
		return nil, err
	}

	// 保存cover文件(放数据库中吧.)
	coverIds, err := NewCreateCoversLogic(
		l.ctx,
		l.svcCtx,
	).CreateCovers(&content.CreateCoversRequest{
		Content: in.Covers,
	})

	// 插入标签，判断标签是否存在如果存在就不添加新的到数据库
	articleId, err := l.svcCtx.ModelWithArticle.InsertReturnId(l.ctx, &model.Article{
		Title:    in.Title,
		Desc:     in.Desc,
		Content:  in.Content,
		Covers:   coverIds.GetIds(),
		Tags:     tagIds.GetIds(),
		Visiable: in.Visiable,
	})
	if err != nil {
		return nil, err
	}

	// 更新标签的中的内容
	modifyTagPushArticleLogic := NewModifyTagPushArticleLogic(l.ctx, l.svcCtx)
	for _, v := range tagIds.GetIds() {
		_, err := modifyTagPushArticleLogic.ModifyTagPushArticle(
			&content.ModifyTagPushArticleRequest{
				Id:      v,
				Article: articleId,
			},
		)
		if err != nil {
			// TODO: 如果对应的tag中添加articleid失败，就添加日志但不处理(一般情况也不会失败).
			logx.Info("tag pusher artilce id warn")
		}
	}

	return &content.CreateArticleResponse{}, nil
}
