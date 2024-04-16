package handler

import (
	"net/http"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetArticleInBackendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetArticleInBackendLogic(r.Context(), svcCtx)
		resp, err := l.GetArticleInBackend()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
