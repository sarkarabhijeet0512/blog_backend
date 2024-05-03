package main

import (
	"blog_api/config"
	"blog_api/internal/server"
	"blog_api/internal/server/handler"
	"blog_api/pkg/cache"
	"blog_api/pkg/posts"
	"blog_api/pkg/rbac"
	"blog_api/pkg/user"
	"blog_api/utils/initialize"

	"go.uber.org/fx"
)

func serverRun() {
	app := fx.New(
		fx.Provide(
			// postgresql
			initialize.NewDB,
			initialize.NewRedisWorker,
		),
		config.Module,
		initialize.Module,
		server.Module,
		handler.Module,
		posts.Module,
		rbac.Module,
		user.Module,
		cache.Module,
	)

	// Run app forever
	app.Run()
}
