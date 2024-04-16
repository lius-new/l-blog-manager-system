package middleware

import (
	"context"
	"net/http"

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
	w.Write([]byte("authorization failed!"))
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		defer func() {
			if err := recover(); err != nil {
				authMiddlewareError(w)
			}
		}()

		secret, err := r.Cookie("secret")
		if err != nil || len(secret.Value) == 0 {
			panic(err)
		}

		judgeResp, err := m.auther.JudgeToken(context.Background(), &authorization.JudgeJwtRequestWithJwt{
			Token: secret.Value,
		})

		if err != nil || len(judgeResp.Id) == 0 {
			panic(err)
		}
		next(w, r)
	}
}
