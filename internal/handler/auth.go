package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/request"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/httperrs"
)

func (h *Handler) signUp(ctx *gin.Context) {
	log := h.logger.Named("signUp")

	var input request.SignUpRequest
	if err := ctx.BindJSON(&input); err != nil {
		httperrs.NewHTTPErrorResponse(ctx, log, err)
		return
	}

	user, err := h.services.Authorization.CreateUser(input.Name, input.Email, input.Password)
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, log, err)
		return
	}
	ctx.JSON(200, user)
}

func (h *Handler) signIn(ctx *gin.Context) {
	log := h.logger.Named("signIn")

	var input request.SignInRequest
	if err := ctx.BindJSON(&input); err != nil {
		httperrs.NewHTTPErrorResponse(ctx, log, err)
		return
	}

	token, err := h.services.Authorization.Authenticate(input.Email, input.Password)
	if err != nil {
		httperrs.NewHTTPErrorResponse(ctx, log, err)
		return
	}
	ctx.JSON(200, map[string]string{
		"token": token,
	})
}
