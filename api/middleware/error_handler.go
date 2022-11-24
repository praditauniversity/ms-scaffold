package middleware

import (
	"ms-scaffold/api/models/web"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ValidationError struct {
	Key     string `json:"key,omitempty"`
	Message string `json:"message"`
}

func (e ValidationError) Error(splitedError []string) []ValidationError {
	var errors []ValidationError
	for _, error := range splitedError {
		splittedError := strings.Split(error, "'")
		errors = append(errors, ValidationError{
			Key:     splittedError[3],
			Message: "Error :" + splittedError[4] + splittedError[5] + splittedError[6],
		})
	}
	return errors
}

type ErrorHandlerMiddleware struct {
	Context *gin.Context
}

func NewErrorHandlerMiddleware(ctx *gin.Context) *ErrorHandlerMiddleware {
	return &ErrorHandlerMiddleware{
		Context: ctx,
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Errors != nil {
			err := c.Errors.Last()
			if err.Meta == "VALIDATION_ERROR" {
				validationErrors(c, err)
				return
			} else if err.Meta == "NOT_FOUND" {
				notFoundError(c, err)
				return
			}
			if err.Meta == "UNAUTHORIZED" {
				authenticationError(c, err)
				return
			}
			internalServerError(c, err)
		}
	}
}

func validationErrors(c *gin.Context, err *gin.Error) {
	splittedError := strings.Split(err.Error(), "\n")
	ValidationError{}.Error(splittedError)
	errorResponse := web.ErrorResponse(http.StatusBadRequest, "VALIDATION_ERROR", ValidationError{}.Error(splittedError))
	c.JSON(http.StatusBadRequest, errorResponse)
}

func authenticationError(c *gin.Context, err *gin.Error) {
	errorResponse := web.ErrorResponse(http.StatusUnauthorized, "UNAUTHORIZED", err)
	c.JSON(http.StatusUnauthorized, errorResponse)
}

func notFoundError(c *gin.Context, err *gin.Error) {
	webResponse := web.ErrorResponse(http.StatusNotFound, "NOT_FOUND", err)
	c.JSON(http.StatusNotFound, webResponse)
}

func internalServerError(c *gin.Context, err *gin.Error) {
	errorResponse := web.ErrorResponse(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err)
	c.JSON(http.StatusInternalServerError, errorResponse)
}