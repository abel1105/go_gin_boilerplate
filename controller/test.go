package controller

import (
	"net/http"

	"github.com/abel1105/go_gin_boilerplate/model"
	"github.com/gin-gonic/gin"
)

type TestController struct{}

var TestModel = &model.Test{}

func (t TestController) Find(ctx *gin.Context) {
	id := ctx.Query("id")

	test, err := TestModel.GetById(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": test,
	})
}

