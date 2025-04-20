package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"redcetarioapi/config"
	"redcetarioapi/controllers"
	"redcetarioapi/middlewares"
	"redcetarioapi/routes"
)

func New(db *gorm.DB, cfg config.Config) *gin.Engine {
	controllers.Init(db)
	r := gin.Default()

	r.Use(middlewares.Logger(cfg.LoggerURL, cfg.LoggerAppKey))

	routes.Setup(r)
	return r
}
