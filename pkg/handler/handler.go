package handler

import (
	"github.com/gin-gonic/gin"
	"user-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.POST("/create", h.create)
		user.GET("/all", h.getAll)
		user.GET("/search", h.getByDateAndAge)
	}

	return router
}
