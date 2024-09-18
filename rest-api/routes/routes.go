package routes

import (
	"github.com/affrianr/auto-reload-waterwind/rest-api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(waterWindHandler *handler.WaterWindHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/waterwind", waterWindHandler.GetAll)
	r.POST("/waterwind", waterWindHandler.Create)
	r.PUT("/waterwind/:id", waterWindHandler.Update)
	r.DELETE("/waterwind/:id", waterWindHandler.Delete)

	return r
}
