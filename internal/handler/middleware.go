package handler

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/internal/entity"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/errs"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/httperrs"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user"
)

func CurrentUser(c *gin.Context) (*entity.User, bool) {
	data, ok := c.Get(userCtx)
	if !ok {
		return nil, false
	}
	acc, ok := data.(*entity.User)
	return acc, ok
}

func MustCurrentUser(c *gin.Context) *entity.User {
	acc, ok := CurrentUser(c)
	if ok {
		return acc
	}
	panic("no user in gin.Context")
}

func (h *Handler) requestIDMiddleware(ctx *gin.Context) {
	ctx.Set("RequestID", uuid.NewV4())
}

func (h *Handler) userIndemnityMiddleware(c *gin.Context) {
	log := h.logger.Named("userIndemnity")

	header := c.GetHeader(authorizationHeader)
	if header == "" {
		httperrs.NewHTTPErrorResponse(c, log, errs.NewError(errs.Unauthorized, "Missing authorization header"))
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		httperrs.NewHTTPErrorResponse(c, log, errs.NewError(errs.Unauthorized, "Invalid authorization header"))
		return
	}

	user, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		httperrs.NewHTTPErrorResponse(c, log, err)
		return
	}

	c.Set(userCtx, &user)
}
