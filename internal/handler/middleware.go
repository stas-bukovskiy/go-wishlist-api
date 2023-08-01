package handler

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func (h *Handler) requestIDMiddleware(ctx *gin.Context) {
	ctx.Set("RequestID", uuid.NewV4())
}
