package middleware

// import (
// 	"fmt"
// 	"ms-scaffold/api/models/web"
// 	"ms-scaffold/api/service"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type AuthorizeJWTMiddleware interface {
// 	AuthorizeJWT() gin.HandlerFunc
// }

// type authorizeJWTMiddleware struct {
// 	jwtService service.JWTService
// }

// func NewAuthorizeJWTMiddleware(jwtService service.JWTService) AuthorizeJWTMiddleware {
// 	return &authorizeJWTMiddleware{
// 		jwtService: jwtService,
// 	}
// }

// func (m *authorizeJWTMiddleware) AuthorizeJWT() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			errorResponse := web.ErrorResponse(http.StatusUnauthorized, "UNAUTHORIZED", "No token provided")
// 			c.JSON(http.StatusUnauthorized, errorResponse)
// 			c.Abort()
// 			return
// 		}
// 		token, err := m.jwtService.ValidateAuthToken(authHeader)
// 		if token.Signature == "invalid" {
// 			errorResponse := web.ErrorResponse(http.StatusUnauthorized, "UNAUTHORIZED", "Invalid token")
// 			c.JSON(http.StatusUnauthorized, errorResponse)
// 			c.Abort()
// 			return
// 		}
// 		if err != nil {
// 			errorResponse := web.ErrorResponse(http.StatusUnauthorized, "UNAUTHORIZED", "Invalid token")
// 			c.JSON(http.StatusUnauthorized, errorResponse)
// 			c.Abort()
// 			return
// 		}
// 		fmt.Println(token.Signature, "ini token signature")
// 	}
// }