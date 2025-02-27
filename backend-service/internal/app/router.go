package app

import (
	question "backend-service/internal/app/domain/question/delivery"
	healthcheck "backend-service/internal/app/domain/system/delivery"
	"context"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router    *gin.Engine
	container *Container
}

func NewRouter(ctx context.Context, router *gin.Engine, container *Container) *Router {
	return &Router{
		router:    router,
		container: container,
	}
}

func (h *Router) RegisterRouter() {

	// h.router.Use(SetTDRMiddleware())
	h.router.Use(gin.Recovery())

	// PING
	h.router.GET("/health", healthcheck.HealthCheck(h.container.HealthCheckService))

	h.router.POST("/questions", question.Ask(h.container.AskService))
}
