package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/lius-new/blog-backend/api"
	"github.com/lius-new/blog-backend/rpc/authorization/auther"
	"github.com/lius-new/blog-backend/rpc/authorization/authorization"
)

type AuthMiddleware struct {
	auther auther.Auther
}

func NewAuthMiddleware(auther auther.Auther) *AuthMiddleware {
	return &AuthMiddleware{auther: auther}
}

func authMiddlewareError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(api.ErrFailedAuthorization.Error()))
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if catchErr := recover(); catchErr != nil {
				authMiddlewareError(w)
			}
		}()

		authorizationString := r.Header.Get("Authorization")

		if len(authorizationString) == 0 || !strings.Contains(authorizationString, "Bearer") {
			panic(api.ErrFailedAuthorization)
		}
		authorizationString = strings.Replace(authorizationString, "Bearer ", "", -1)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		judgeResp, err := m.auther.JudgeToken(ctx, &authorization.JudgeJwtRequestWithJwt{
			Token: authorizationString,
		})

		if err != nil || len(judgeResp.Id) == 0 {
			panic(err)
		}
		next(w, r)
	}
}
