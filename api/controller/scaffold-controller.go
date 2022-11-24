package controller

import (
	"ms-scaffold/api/exception"
	"ms-scaffold/api/helper"
	"ms-scaffold/api/models/web"
	"ms-scaffold/api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var scaffoldLog = "scaffold.log"

type ScaffoldController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type scaffoldController struct {
	scaffoldService service.ScaffoldService
	jwtService      service.JWTService
	logger          helper.Log
}

func NewScaffoldController(scaffoldService service.ScaffoldService /*, jwtService service.JWTService*/, logger helper.Log) ScaffoldController {
	return &scaffoldController{
		scaffoldService: scaffoldService,
		// jwtService:      jwtService,
		logger: logger,
	}
}

func (controller *scaffoldController) All(context *gin.Context) {
	controller.logger.SetUp(scaffoldLog)
	scaffold := controller.scaffoldService.All()
	webResponse := web.SuccessResponse(scaffold)
	context.JSON(http.StatusOK, webResponse)
	// token := context.GetHeader("Authorization")
	// userId, _ := controller.jwtService.GetUserData(token, "user_id")
	// controller.logger.Infof("%s already get all scaffolds", userId)
}

func (controller *scaffoldController) FindById(context *gin.Context) {
	controller.logger.SetUp(scaffoldLog)
	notFoundError := exception.NewNotFoundError(context, controller.logger)
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	scaffold, err := controller.scaffoldService.FindById(uint(id))
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	webResponse := web.SuccessResponse(scaffold)
	context.JSON(http.StatusOK, webResponse)
	// token := context.GetHeader("Authorization")
	// userId, _ := controller.jwtService.GetUserData(token, "user_id")
	// controller.logger.Infof("%s already find a user data with id %d", userId, scaffold.ID)
}

func (controller *scaffoldController) Create(context *gin.Context) {
	controller.logger.SetUp(scaffoldLog)
	validationError := exception.NewValidationError(context, controller.logger)
	internalServerError := exception.NewInternalServerError(context, controller.logger)
	var request web.ScaffoldRequest
	err := context.BindJSON(&request)
	ok := validationError.SetMeta(err)
	if ok {
		validationError.Logf(err)
		return
	}
	scaffold, err := controller.scaffoldService.Create(request)

	ok = internalServerError.SetMeta(err)
	if ok {
		internalServerError.Logf(err)
		return
	}

	webResponse := web.SuccessResponse(scaffold)
	context.JSON(http.StatusCreated, webResponse)
	// token := context.GetHeader("Authorization")
	// userId, _ := controller.jwtService.GetUserData(token, "user_id")
	// controller.logger.Infof("%d already insert a scaffold with id %d", userId, scaffold.ID)
}

func (controller *scaffoldController) Update(context *gin.Context) {
	controller.logger.SetUp(scaffoldLog)
	validatioError := exception.NewValidationError(context, controller.logger)
	notFoundError := exception.NewNotFoundError(context, controller.logger)
	var request web.ScaffoldUpdateRequest
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	request.ID = uint(id)
	err = context.BindJSON(&request)
	ok = validatioError.SetMeta(err)
	if ok {
		validatioError.Logf(err)
		return
	}
	scaffold, err := controller.scaffoldService.Update(request)
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	webResponse := web.SuccessResponse(scaffold)
	context.JSON(http.StatusOK, webResponse)
	// token := context.GetHeader("Authorization")
	// userId, _ := controller.jwtService.GetUserData(token, "user_id")
	// controller.logger.Infof("%s already updated a scaffold with id %d", userId, scaffold.ID)
}

func (controller *scaffoldController) Delete(context *gin.Context) {
	controller.logger.SetUp(scaffoldLog)
	notFoundError := exception.NewNotFoundError(context, controller.logger)
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	err = controller.scaffoldService.Delete(uint(id))
	ok = notFoundError.SetMeta(err)
	if ok {
		notFoundError.Logf(err)
		return
	}
	webResponse := web.SuccessResponse(nil)
	context.JSON(http.StatusOK, webResponse)
	// token := context.GetHeader("Authorization")
	// userId, _ := controller.jwtService.GetUserData(token, "user_id")
	// controller.logger.Infof("%s already deleted a scaffold with id %d", userId, id)
}
