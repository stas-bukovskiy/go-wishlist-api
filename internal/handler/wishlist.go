package handler

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/request"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/errs"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/httperrs"
	"net/http"
)

// @Summary      List all wishlists
// @Description  Get all wishlists of this user
// @Security 	 ApiKeyAuth
// @Tags         wishlists
// @Accept       json
// @Produce      json
// @success      200 {array} entity.Wishlist "wishlists"
// @Failure      400,500  {object}  httperrs.ErrorResponse
// @Router       /api/v1/wishlists/ [get]
func (h *Handler) getAllWishlists(ctx *gin.Context) {
	user := MustCurrentUser(ctx)

	wishlists, err := h.services.Wishlist.GetAllByUserID(user.ID)
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, err)
		return
	}
	ctx.JSON(http.StatusOK, wishlists)
}

// @Summary      Get wishlist
// @Description  Get wishlist by its id
// @Security 	 ApiKeyAuth
// @Tags         wishlists
// @Accept       json
// @Produce      json
// @Param 		 id path string true "wishlist id"
// @success      200 {object} entity.Wishlist "wishlist"
// @Failure      400,404,500  {object}  httperrs.ErrorResponse
// @Router       /api/v1/wishlists/{id} [get]
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

// @Summary      Get wishlist items
// @Description  Get wishlist items by its id
// @Security 	 ApiKeyAuth
// @Tags         wishlists
// @Accept       json
// @Produce      json
// @Param 		 id path string true "wishlist id"
// @success      200 {array} entity.WishlistItem "wishlist items"
// @Failure      400,404,500  {object}  httperrs.ErrorResponse
// @Router       /api/v1/wishlists/{id}/items [get]
func (h *Handler) getWishlistItemsByID(ctx *gin.Context) {
	wishlistID, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, errs.NewError(errs.Validation, "wishlist id is not valid"))
		return
	}

	wishlist, err := h.services.Wishlist.GetItemsByID(wishlistID)
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, err)
		return
	}
	ctx.JSON(http.StatusOK, wishlist)
}

// @Summary      Create a new wishlist
// @Description  Create a new, empty wishlist
// @Security 	 ApiKeyAuth
// @Tags         wishlists
// @Accept       json
// @Produce      json
// @Param 		 request body request.WishlistRequest true "wishlist request"
// @success      200 {object} entity.Wishlist "created wishlist"
// @Failure      400,500  {object}  httperrs.ErrorResponse
// @Router       /api/v1/wishlists/ [post]
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

// @Summary      Update a wishlist
// @Description  Update a wishlist by its id
// @Security 	 ApiKeyAuth
// @Tags         wishlists
// @Accept       json
// @Produce      json
// @Param 		 id path string true "wishlist id"
// @Param 		 request body request.WishlistRequest true "wishlist request"
// @success      200 {object} entity.Wishlist "updated wishlist"
// @Failure      400,404,500  {object}  httperrs.ErrorResponse
// @Router       /api/v1/wishlists/{id} [put]
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

// @Summary      Delete a wishlist
// @Description  Delete a wishlist by its id
// @Security 	 ApiKeyAuth
// @Tags         wishlists
// @Accept       json
// @Produce      json
// @Param 		 id path string true "wishlist id"
// @success      200 {object} entity.Wishlist "deleted wishlist"
// @Failure      400,404,500  {object}  httperrs.ErrorResponse
// @Router       /api/v1/wishlists/{id} [delete]
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
