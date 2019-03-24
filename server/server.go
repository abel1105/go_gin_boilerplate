package server

import (
	"github.com/abel1105/go_gin_boilerplate/config"
	"github.com/abel1105/go_gin_boilerplate/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {
	env := config.GetConfig().GetString("env")
	port := config.GetConfig().GetString("port")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()
	r.Use(middleware.HandleErrors())
	r = NewRouter(r)
	_ = r.Run(":" + port)
}