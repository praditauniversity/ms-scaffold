package routes

import (
	"ms-scaffold/api/injector"
	"ms-scaffold/api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewScaffoldRoutes(db *gorm.DB, route *gin.Engine) {
	scaffoldController := injector.InitScaffold(db)
	scaffoldRoute := route.Group("/api/v1/scaffold")
	scaffoldRoute.Use(middleware.ErrorHandler())
	scaffoldRoute.Use(cors.Default())
	scaffoldRoute.GET("/", scaffoldController.All)
	scaffoldRoute.GET("/:id", scaffoldController.FindById)
	scaffoldRoute.POST("/", scaffoldController.Insert)
	scaffoldRoute.PUT("/:id", scaffoldController.Update)
	scaffoldRoute.DELETE("/:id", scaffoldController.Delete)
}
