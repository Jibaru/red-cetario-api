package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"redcetarioapi/controllers"
	"redcetarioapi/routes"
)

func New(db *gorm.DB) *gin.Engine {
	controllers.Init(db)
	r := gin.Default()
	routes.Setup(r)
	return r
}
