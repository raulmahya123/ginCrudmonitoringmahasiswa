package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func MataKuliahRoute(router *gin.Engine) {
	router.POST("/matakuliah", controllers.CreateMatakuliah())
	router.GET("/matakuliah/:matakuliahGetID", controllers.GetMatakuliah())
	router.PUT("/matakuliah/:matakuliahID", controllers.EditMatakuliah())
	router.DELETE("/matakuliah/:matakuliahID", controllers.DeleteMatakuliah())
	router.GET("/matakuliahs", controllers.GetAllMatakuliahs())
}
