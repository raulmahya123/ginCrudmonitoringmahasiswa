package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func OrangTuaRoute(router *gin.Engine) {
	router.POST("/orangtua", controllers.CreateOrangtua())
	router.GET("/orangtua/:orangtuaGetID", controllers.GetOrangtua())
	router.PUT("/orangtua/:orangtuaID", controllers.EditOrangtua())
	router.DELETE("/orangtua/:orangtuaID", controllers.DeleteOrangtua())
	router.GET("/orangtuas", controllers.GetAllOrangtuas())
}
