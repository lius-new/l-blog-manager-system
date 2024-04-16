package handler

import (
	"net/http"

	"github.com/lius-new/blog-backend/api"
	"github.com/lius-new/blog-backend/api/user/internal/logic"
	"github.com/lius-new/blog-backend/api/user/internal/svc"
	"github.com/lius-new/blog-backend/api/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ModifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewModifyLogic(r.Context(), svcCtx)
		resp, err := l.Modify(&req)
		api.Response(w, resp, err)

	}
}
