package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/service"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
)

type Handler struct {
	services *service.Service
	logger   logger.Logger
}

func NewHandler(services *service.Service, logger logger.Logger) *Handler {
	return &Handler{services: services, logger: logger}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(h.requestIDMiddleware)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api/v1", h.userIndemnityMiddleware)
	{
		wishlists := api.Group("/wishlists")
		{
			wishlists.GET("/", h.getAllWishlists)
			wishlists.GET("/:id", h.getWishlistByID)
			wishlists.POST("/", h.createWishlist)
			wishlists.PUT("/:id", h.updateWishlist)
			wishlists.DELETE("/:id", h.deleteWishlist)
		}
		wishlistItems := api.Group("/wishlist-items")
		{
			wishlistItems.GET("/:id", h.getWishlistItem)
			wishlistItems.POST("/", h.addItemToWishlist)
			wishlistItems.PUT("/:id", h.updateItem)
			wishlistItems.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
