package handler

import (
	"net/http"

	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateArticlesInBackendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqInBackend
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateArticlesInBackendLogic(r.Context(), svcCtx)
		resp, err := l.CreateArticlesInBackend(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
