package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/request"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/httperrs"
)

// @Summary      Sign up
// @Description  Create user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request    body     request.SignUpRequest  true  "body request to sign up"
// @Success      200  {object}  entity.User
// @Failure      400,500  {object}  httperrs.ErrorResponse
// @Router       /auth/sign-up [post]
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

// @Summary      Sign in
// @Description  Create user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        q    body     request.SignInRequest  true  "body request to sign in"
// @success      200 {object} map[string]string{} "access token"
// @Failure      400,500  {object}  httperrs.ErrorResponse
// @Router       /auth/sign-in [post]
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
