package httperrs

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/errs"
	"github.com/stas-bukovskiy/go-n-react-wishlist-app/pkg/logger"
	"net/http"
)

type ErrorResponse struct {
	Error ServiceError `json:"error"`
}

type ServiceError struct {
	Kind    string `json:"kind,omitempty"`
	Param   string `json:"param,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewHTTPErrorResponse(ctx *gin.Context, log logger.Logger, err error) {
	if err == nil {
		nilHTTPErrorResponse(ctx, log)
		return
	}

	var e *errs.Error
	if errors.As(err, &e) {
		typicalHTTPErrorResponse(ctx, log, e)
		return
	}

	unknownHTTPErrorResponse(ctx, log, err)
}

func nilHTTPErrorResponse(ctx *gin.Context, log logger.Logger) {
	log.WithContext(ctx).Error("nil error - no response body sent")
	ctx.AbortWithStatus(http.StatusInternalServerError)
}

func typicalHTTPErrorResponse(ctx *gin.Context, log logger.Logger, e *errs.Error) {
	httpStatusCode := httpStatusCode(e.Kind)

	if e.IsEmpty() {
		log.WithContext(ctx).Error("error sent but with empty body")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if httpStatusCode == http.StatusInternalServerError {
		errorLogger(ctx, log, httpStatusCode, e).Error("unexpected error: %+v", e)
	} else {
		errorLogger(ctx, log, httpStatusCode, e).Debug("error response sent to client")
	}

	errorResponse := newErrResponse(e)
	ctx.AbortWithStatusJSON(httpStatusCode, errorResponse)
}

func errorLogger(ctx *gin.Context, log logger.Logger, httpStatusCode int, e *errs.Error) logger.Logger {
	if ops := errs.OpStack(e); len(ops) > 0 {
		return log.WithContext(ctx).With(map[string]interface{}{
			"httpStatusCode": httpStatusCode,
			"kind":           e.Kind.String(),
			"parameter":      e.Parameter,
			"stacktrace":     fmt.Sprintf("%+v", ops),
		})
	} else {
		return log.With(map[string]interface{}{
			"httpStatusCode": httpStatusCode,
			"kind":           e.Kind.String(),
			"parameter":      e.Parameter,
		})
	}
}

func unknownHTTPErrorResponse(ctx *gin.Context, log logger.Logger, e error) {
	er := ErrorResponse{
		Error: ServiceError{
			Message: "Unexpected error - contact support",
		},
	}
	log.WithContext(ctx).Error("Unexpected error: %+v", e)
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, er)
}

func newErrResponse(e *errs.Error) ErrorResponse {
	const msg string = "internal server error - please contact support"

	switch e.Kind {
	case errs.Internal, errs.Database:
		return ErrorResponse{
			Error: ServiceError{
				Kind:    errs.Internal.String(),
				Message: msg,
			},
		}
	default:
		return ErrorResponse{
			Error: ServiceError{
				Kind:    e.Kind.String(),
				Param:   string(e.Parameter),
				Message: e.Error(),
			},
		}
	}
}

func httpStatusCode(kind errs.Kind) int {
	switch kind {
	case errs.IO, errs.Internal, errs.Database:
		return http.StatusInternalServerError
	case errs.AlreadyExist:
		return http.StatusConflict
	case errs.NotFound:
		return http.StatusNotFound
	case errs.Validation, errs.BadRequest:
		return http.StatusBadRequest
	case errs.Unauthorized:
		return http.StatusUnauthorized
	case errs.Unauthenticated:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
