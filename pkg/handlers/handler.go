package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mihailco/memessenger/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{services: serv}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.GET("/messages")
		// lists := api.Group("/lists")
		// {
		// 	lists.POST("/")
		// 	lists.GET("/")
		// 	lists.GET("/:id")
		// 	lists.PUT("/:id")
		// 	lists.DELETE("/:id")
		// 	items := router.Group(":id/items")
		// 	{
		// 		items.POST("/")
		// 		items.GET("/")
		// 		items.GET("/:item_id")

		// 	}
		// }
	}
	return router
}
