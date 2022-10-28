package controller

import (
	"ms-scaffold/api/helper"
	"ms-scaffold/api/models/web"
	"ms-scaffold/api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ScaffoldController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type scaffoldController struct {
	scaffoldService service.ScaffoldService
}

func NewScaffoldController(scaffoldService service.ScaffoldService) ScaffoldController {
	return &scaffoldController{
		scaffoldService: scaffoldService,
	}
}

func (controller *scaffoldController) All(context *gin.Context) {
	scaffold := controller.scaffoldService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   scaffold,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (controller *scaffoldController) FindById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	scaffold, err := controller.scaffoldService.FindById(uint(id))
	ok = helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   scaffold,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (controller *scaffoldController) Insert(context *gin.Context) {
	var request web.ScaffoldRequest
	err := context.BindJSON(&request)
	ok := helper.InternalServerError(context, err)
	if ok {
		return
	}
	scaffold, err := controller.scaffoldService.Create(request)

	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   scaffold,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (controller *scaffoldController) Update(context *gin.Context) {
	var request web.ScaffoldUpdateRequest
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	request.ID = uint(id)
	err = context.BindJSON(&request)
	ok = helper.ValidationError(context, err)
	if ok {
		return
	}
	scaffold, err := controller.scaffoldService.Update(request)
	ok = helper.InternalServerError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   scaffold,
	}
	context.JSON(http.StatusOK, webResponse)
}

func (controller *scaffoldController) Delete(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	ok := helper.NotFoundError(context, err)
	if ok {
		return
	}
	err = controller.scaffoldService.Delete(uint(id))
	ok = helper.NotFoundError(context, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: "",
		Data:   "Scaffold has been removed",
	}
	context.JSON(http.StatusOK, webResponse)
}
