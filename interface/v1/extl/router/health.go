package router

import (
	"github.com/labstack/echo/v4"

	"codebase/app/v1/deps"
	"codebase/interface/v1/extl/handler/health"
)

func healthRouter(e *echo.Group, deps deps.IDependency) {
	base := deps.GetBase()
	handler := health.New(base.Mdbc, base.Rdbc)
	e.GET("/health", handler.Health)
}
