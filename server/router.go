package server

import (
	"github.com/abel1105/go_gin_boilerplate/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine) *gin.Engine {
	v1 := router.Group("/v1")
	{
		test := &controller.TestController{}
		v1.GET("/find", test.Find)
	}

	return router
}