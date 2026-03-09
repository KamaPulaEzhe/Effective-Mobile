package handler

import (
	"github.com/effective/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		subs := api.Group("/subscriptions")
		{
			subs.GET("/total-cost", h.getCost)
			subs.POST("/", h.createSub)
			subs.GET("/:id", h.getSub)
			subs.PATCH("/:id", h.updateSub)
			subs.DELETE("/:id", h.deleteSub)
		}
	}

	return router
}
