package handler

import (
	"net/http"

	"github.com/lius-new/blog-backend/api"
	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetArticlesByPageWithViewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetArticleByPageWithViewRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetArticlesByPageWithViewLogic(r.Context(), svcCtx)
		resp, err := l.GetArticlesByPageWithView(&req)
		api.Response(w, resp, err)

	}
}