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
