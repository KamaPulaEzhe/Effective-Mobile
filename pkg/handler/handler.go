package handler

import (
	_ "github.com/effective/docs"
	"github.com/effective/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
