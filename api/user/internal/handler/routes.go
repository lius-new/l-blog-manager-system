// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/lius-new/blog-backend/api/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/users"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/",
					Handler: ModifyHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: SelectByPageHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id",
					Handler: DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: SelectHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/articles"),
	)
}