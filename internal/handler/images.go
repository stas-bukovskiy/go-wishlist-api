package handler

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/validation"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/errs"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/httperrs"
	"net/http"
)

func (h *Handler) uploadImage(ctx *gin.Context) {
	file, fileHeaders, err := ctx.Request.FormFile("image")
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, errs.NewError(errs.BadRequest, err))
		return
	}

	if !validation.IsValidImageName(fileHeaders.Filename) {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, errs.NewError(errs.Validation, "invalid image filename: should ends with png, jpeg or jpg"))
		return
	}

	image, err := h.services.Image.CreateImage(file, fileHeaders)
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, err)
		return
	}
	ctx.JSON(http.StatusOK, image)
}

func (h *Handler) deleteImage(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, errs.NewError(errs.Validation, "invalid image id"))
	}

	err = h.services.Image.DeleteImage(id)
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, h.logger, err)
		return
	}
	ctx.Status(http.StatusOK)
}
