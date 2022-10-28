package helper

import "github.com/gin-gonic/gin"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func NotFoundError(ctx *gin.Context, err error) bool {
	if err != nil {
		ctx.Error(err).SetMeta("NOT_FOUND")
		return true
	}
	return false
}

func ValidationError(ctx *gin.Context, err error) bool {
	if err != nil {
		ctx.Error(err).SetMeta("VALIDATION_ERROR")
		return true
	}
	return false
}

func InternalServerError(ctx *gin.Context, err error) bool {
	if err != nil {
		ctx.Error(err).SetMeta("INTERNAL_SERVER_ERROR")
		return true
	}
	return false
}

func AuthenticationError(ctx *gin.Context, err error) bool {
	if err != nil {
		ctx.Error(err).SetMeta("UNAUTHORIZED")
		return true
	}
	return false
}

func TokenError(ctx *gin.Context, err error) bool {
	if err != nil {
		ctx.Error(err).SetMeta("TOKEN_ERROR")
		return true
	}
	return false
}
