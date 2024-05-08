package handler

import (
	"net/http"

	"github.com/lius-new/blog-backend/api"
	"github.com/lius-new/blog-backend/api/user/internal/logic"
	"github.com/lius-new/blog-backend/api/user/internal/svc"
	"github.com/lius-new/blog-backend/api/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SelectByPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SelectPage
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSelectByPageLogic(r.Context(), svcCtx)
		resp, err := l.SelectByPage(&req)
		api.Response(w, resp, err)
	}
}
