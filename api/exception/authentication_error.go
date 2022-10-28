package exception

import (
	"ms-scaffold/api/helper"

	"github.com/gin-gonic/gin"
)

type AuthenticationError struct {
	context *gin.Context
	log     helper.Log
}

func NewAuthenticationError(context *gin.Context, log helper.Log) AuthenticationError {
	return AuthenticationError{
		context: context,
		log:     log,
	}
}

// this code is used to set the meta of the error
func (tokenError AuthenticationError) SetMeta(message error) bool {
	if message != nil {
		tokenError.context.Error(message).SetMeta("UNAUTHORIZED")
		tokenError.context.Status(401)
		return true
	}
	return false
}

// this code is used to log the error
func (tokenError AuthenticationError) Logf(message error, args ...interface{}) {
	tokenError.log.Errorf("Authentication Error : "+message.Error(), args...)
}
