//go:build wireinject
// +build wireinject

package injector

import (
	"ms-scaffold/api/controller"
	"ms-scaffold/api/repository"
	"ms-scaffold/api/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var scaffoldSet = wire.NewSet(
	repository.NewScaffoldRepository,
	service.NewScaffoldService,
	controller.NewScaffoldController,
)

func InitScaffold(db *gorm.DB) controller.ScaffoldController {
	wire.Build(
		scaffoldSet,
	)
	return nil
}
