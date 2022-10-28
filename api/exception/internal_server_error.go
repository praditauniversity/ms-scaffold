package exception

import (
	"ms-scaffold/api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InternalServerError struct {
	context *gin.Context
	log     helper.Log
}

func NewInternalServerError(context *gin.Context, log helper.Log) InternalServerError {
	return InternalServerError{
		context: context,
		log:     log,
	}
}

// this code is used to set the meta of the error
func (e InternalServerError) SetMeta(message error) bool {
	if message != nil {
		e.context.Error(message).SetMeta("INTERNAL_SERVER_ERROR")
		e.context.Status(http.StatusInternalServerError)
		return true
	}
	return false
}

// this code is used to log the error
func (e InternalServerError) Logf(message error, args ...interface{}) {
	e.log.Errorf("Internal Server Error : "+message.Error(), args...)
}
