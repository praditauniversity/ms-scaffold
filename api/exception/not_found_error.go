package exception

import (
	"ms-scaffold/api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotFoundError struct {
	context *gin.Context
	log     helper.Log
}

func NewNotFoundError(context *gin.Context, log helper.Log) NotFoundError {
	return NotFoundError{
		context: context,
		log:     log,
	}
}

// this code is used to set the meta of the error
func (tokenError NotFoundError) SetMeta(message error) bool {
	if message != nil {
		tokenError.context.Error(message).SetMeta("NOT_FOUND")
		tokenError.context.Status(http.StatusNotFound)
		return true
	}
	return false
}

// this code is used to log the error
func (tokenError NotFoundError) Logf(message error, args ...interface{}) {
	tokenError.log.Errorf("Not Found Error : "+message.Error(), args...)
}
