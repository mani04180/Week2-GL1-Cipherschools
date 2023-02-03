package routers

import (
	handler "https://github.com/mani04180/Handler"
	"https://github.com/mani04180/database"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	api := handler.Handler{
		DB: database.GetDB(),
	}
	router.GET("/Books", api.GetBooks)
	router.POST("/Books", api.SaveBook)
	return router
}
