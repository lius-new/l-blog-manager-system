package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/lius-new/blog-backend/api"
	"github.com/lius-new/blog-backend/api/article/internal/logic"
	"github.com/lius-new/blog-backend/api/article/internal/svc"
	"github.com/lius-new/blog-backend/api/article/internal/types"
)

func ViewImageWithBase64Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ViewImageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewViewImageLogic(r.Context(), svcCtx)
		resp, err := l.ViewImage(&req)

		errHandle := func(err error) {
			var body api.Body
			body.Code = -1
			body.Msg = err.Error()
			httpx.OkJson(w, body)
		}

		if err != nil {
			errHandle(err)
		} else if resp == nil {
			errHandle(api.ErrInvalidNotFound)
		} else {
				w.Write([]byte(resp.Base64))
				httpx.Ok(w)
		}
	}
}
