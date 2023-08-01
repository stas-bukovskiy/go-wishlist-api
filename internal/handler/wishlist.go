package handler

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/request"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/errs"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/httperrs"
	"net/http"
)

func (h *Handler) getAllWishlists(ctx *gin.Context) {
	user := MustCurrentUser(ctx)

	wishlists, err := h.services.Wishlist.GetAllByUserID(user.ID)
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, err)
		return
	}
	ctx.JSON(http.StatusOK, wishlists)
}

func (h *Handler) getWishlistByID(ctx *gin.Context) {
	wishlistID, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, errs.NewError(errs.Validation, "wishlist id is not valid"))
		return
	}

	wishlist, err := h.services.Wishlist.GetByID(wishlistID)
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, err)
		return
	}
	ctx.JSON(http.StatusOK, wishlist)
}

func (h *Handler) createWishlist(ctx *gin.Context) {
	user := MustCurrentUser(ctx)

	var input request.WishlistRequest
	if err := ctx.BindJSON(&input); err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, errs.NewError(errs.Validation, err))
		return
	}

	wishlist, err := h.services.Wishlist.CreateWishlist(user.ID, input.Title, input.Description)
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, err)
		return
	}
	ctx.JSON(http.StatusCreated, wishlist)
}

func (h *Handler) updateWishlist(ctx *gin.Context) {
	wishlistID, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, errs.NewError(errs.Validation, "wishlist id is not valid"))
		return
	}

	var input request.WishlistRequest
	if err := ctx.BindJSON(&input); err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, errs.NewError(errs.Validation, err))
		return
	}

	wishlist, err := h.services.Wishlist.UpdateWishlist(wishlistID, input.Title, input.Description)
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, err)
		return
	}
	ctx.JSON(http.StatusOK, wishlist)
}

func (h *Handler) deleteWishlist(ctx *gin.Context) {
	wishlistID, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, errs.NewError(errs.Validation, "wishlist id is not valid"))
		return
	}

	wishlist, err := h.services.Wishlist.DeleteWishlist(wishlistID)
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, err)
		return
	}
	ctx.JSON(http.StatusOK, wishlist)
}
