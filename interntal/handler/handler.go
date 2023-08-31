package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/naumovrus/backend-trainee-asignment/interntal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api", gin.BasicAuth(gin.Accounts{
		"admin": "qwerty",
	}))
	{
		user := api.Group("/user")
		{
			user.POST("/", h.createUser)
			user.GET("/:userId", h.getUserSegments)
		}
		segments := api.Group("segments")
		{
			segments.POST("/", h.createSegment)
			segments.POST("/:userId", h.addUserSegment)
			segments.DELETE("/:userId", h.deleteUserSegment)
			segments.DELETE("/", h.deleteSegment)
		}

	}
	return router
}
