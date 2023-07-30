package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/service"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
)

type Handler struct {
	services *service.Service
	logger   *logger.Logger
}

func NewHandler(services *service.Service, logger *logger.Logger) *Handler {
	return &Handler{services: services, logger: logger}
}

func (h *Handler) InitRoutes() {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}
}
