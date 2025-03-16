package app

import (
	knowledge "backend-service/internal/app/domain/knowledge/delivery"
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

	limiter := &RateLimiter{}

	h.router.Use(SetTDRMiddleware())
	h.router.Use(gin.Recovery())

	// PING
	h.router.GET("/health", healthcheck.HealthCheck(h.container.HealthCheckUsecase))

	// API Routes
	api := h.router.Group("/api")
	{
		// V1 Routes
		v1 := api.Group("/v1")
		{
			v1.POST("/ask", limiter.RateLimitMiddleware(), question.Ask(h.container.AskUsecase))
			v1.POST("/store", StaticKeyMiddleware(), knowledge.Store(h.container.StoreUsecase))
		}

		// V2 Routes
		v2 := api.Group("/v2")
		{
			v2.POST("/ask", limiter.RateLimitMiddleware(), question.Ask(h.container.AskUsecase))
			v2.POST("/store", StaticKeyMiddleware(), knowledge.StoreV2(h.container.StoreUsecase))
		}
	}
}
