package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func MahasiswaRoute(router *gin.Engine) {
	router.POST("/mahasiswa", controllers.CreateMahasiswa())
	router.GET("/mahasiswa/:mahasiswaGetid", controllers.GetMahasiswa())
	router.PUT("/mahasiswa/:mahasiswaID", controllers.EditMahasiswa())
	router.DELETE("/mahasiswa/:mahasiswaID", controllers.DeleteMahasiswa())
	router.GET("/mahasiswas", controllers.GetAllMahasiswas())
}
