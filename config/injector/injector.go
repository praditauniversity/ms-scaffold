//go:build wireinject
// +build wireinject

package injector

import (
	"ms-scaffold/api/controller"
	"ms-scaffold/api/helper"
	"ms-scaffold/api/middleware"
	"ms-scaffold/api/repository"
	"ms-scaffold/api/service"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var jwtMiddlewareSet = wire.NewSet(
	service.NewJWTService,
	middleware.NewAuthorizeJWTMiddleware,
)

var scaffoldSet = wire.NewSet(
	logrus.New,
	helper.NewLog,
	controller.NewScaffoldController,
	repository.NewScaffoldRepository,
	service.NewJWTService,
	service.NewScaffoldService,
)

func InitScaffold(db *gorm.DB) controller.ScaffoldController {
	wire.Build(
		scaffoldSet,
	)
	return nil
}

func InitJWTMiddleware() middleware.AuthorizeJWTMiddleware {
	wire.Build(
		jwtMiddlewareSet,
	)
	return nil
}
