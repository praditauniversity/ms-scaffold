package exception

import (
	"ms-scaffold/api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ValidationError struct {
	context *gin.Context
	log     helper.Log
}

func NewValidationError(context *gin.Context, log helper.Log) ValidationError {
	return ValidationError{
		context: context,
		log:     log,
	}
}

// this code is used to set the meta of the error
func (e ValidationError) SetMeta(message error) bool {
	if message != nil {
		e.context.Error(message).SetMeta("VALIDATION_ERROR")
		e.context.Status(http.StatusBadRequest)
		return true
	}
	return false
}

// this code is used to log the error
func (e ValidationError) Logf(message error, args ...interface{}) {
	e.log.Errorf("Validation Error : "+message.Error(), args...)
}
