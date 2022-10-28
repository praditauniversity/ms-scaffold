package middleware

import (
	"ms-scaffold/api/models/web"
	"ms-scaffold/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorizeJWTMiddleware interface {
	AuthorizeJWT() gin.HandlerFunc
}

type authorizeJWTMiddleware struct {
	jwtService service.JWTService
}

func NewAuthorizeJWTMiddleware(jwtService service.JWTService) AuthorizeJWTMiddleware {
	return &authorizeJWTMiddleware{
		jwtService: jwtService,
	}
}

func (m *authorizeJWTMiddleware) AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Errors: "Not token found",
				Data:   nil,
			}
			c.JSON(http.StatusUnauthorized, webResponse)
			c.Abort()
			return
		}
		_, err := m.jwtService.ValidateToken(authHeader)
		if err != nil {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
				Errors: "Token is invalid",
				Data:   nil,
			}
			c.JSON(http.StatusUnauthorized, webResponse)
			c.Abort()
			return
		}
	}
}
