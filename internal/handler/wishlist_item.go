package handler

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/request"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/errs"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/httperrs"
	"net/http"
)

// @Summary      Get wishlist item
// @Description  Get wishlist item by its id
// @Security 	 ApiKeyAuth
// @Tags         wishlist-items
// @Accept       json
// @Produce      json
// @Param 		 id path string true "wishlist item id"
// @success      200 {object} entity.WishlistItem "wishlist item"
// @Failure      400,404,500  {object}  httperrs.ErrorResponse
// @Router       /api/v1/wishlist-items/{id} [get]
func (h *Handler) getWishlistItem(ctx *gin.Context) {
	itemID, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, errs.NewError(errs.Validation, "item id is invalid"))
		return
	}

	item, err := h.services.WishlistItem.GetByID(itemID)
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, err)
		return
	}
	ctx.JSON(http.StatusOK, item)
}

// @Summary      Create a new wishlist item
// @Description  Create a new wishlist item
// @Security 	 ApiKeyAuth
// @Tags         wishlist-items
// @Accept       json
// @Produce      json
// @Param 		 request body request.AddWishlistItemRequest true "create wishlist item request"
// @success      200 {object} entity.WishlistItem "wishlist item"
// @Failure      400,404,500  {object}  httperrs.ErrorResponse
// @Router       /api/v1/wishlist-items/ [post]
func (h *Handler) addItemToWishlist(ctx *gin.Context) {
	var itemToAdd request.AddWishlistItemRequest
	if err := ctx.BindJSON(&itemToAdd); err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, errs.NewError(errs.Validation, err))
		return
	}

	item, err := h.services.WishlistItem.AddItemToWishlist(entity.WishlistItem{
		WishlistId:  itemToAdd.WishlistId,
		Title:       itemToAdd.Title,
		Description: itemToAdd.Description,
		Price:       itemToAdd.Price,
		Images:      fromIdToImages(itemToAdd.ImageIDs),
		URL:         itemToAdd.URL,
	})
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, err)
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func fromIdToImages(imageIDs []uuid.UUID) []entity.Image {
	var images []entity.Image
	for _, imageId := range imageIDs {
		images = append(images, entity.Image{
			Base: entity.Base{
				ID: imageId,
			},
		})
	}
	return images
}

// @Summary      Update a wishlist item
// @Description  Create existing wishlist item by its id
// @Security 	 ApiKeyAuth
// @Tags         wishlist-items
// @Accept       json
// @Produce      json
// @Param 		 id path string true "wishlist item id"
// @Param 		 request body request.UpdateWishlistItemRequest true "update wishlist item request"
// @success      200 {object} entity.WishlistItem "wishlist item"
// @Failure      400,404,500  {object}  httperrs.ErrorResponse
// @Router       /api/v1/wishlist-items/{id} [put]
func (h *Handler) updateItem(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, errs.NewError(errs.Validation, "invalid item id"))
		return
	}

	var itemToUpdate request.UpdateWishlistItemRequest
	if err := ctx.BindJSON(&itemToUpdate); err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, errs.NewError(errs.Validation, err))
		return
	}

	item, err := h.services.WishlistItem.UpdateItem(id, entity.WishlistItem{
		Title:       itemToUpdate.Title,
		Description: itemToUpdate.Description,
		Price:       itemToUpdate.Price,
		Images:      fromIdToImages(itemToUpdate.ImageIDs),
	})

	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, err)
		return
	}
	ctx.JSON(http.StatusOK, item)
}

// @Summary      Delete wishlist item
// @Description  Delete wishlist item by its id
// @Security 	 ApiKeyAuth
// @Tags         wishlist-items
// @Accept       json
// @Produce      json
// @Param 		 id path string true "wishlist item id"
// @success      200 {object} entity.WishlistItem "wishlist item"
// @Failure      400,404,500  {object}  httperrs.ErrorResponse
// @Router       /api/v1/wishlist-items/{id} [delete]
func (h *Handler) deleteItem(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, errs.NewError(errs.Validation, "invalid item id"))
		return
	}

	item, err := h.services.WishlistItem.DeleteItem(id)

	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, err)
		return
	}
	ctx.JSON(http.StatusOK, item)
}
