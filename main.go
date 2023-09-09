package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes" //add this

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.AbsensiRoute(router)   //add this
	routes.OrangTuaRoute(router)  //add this
	routes.NilaiRoute(router)     //add this
	routes.MahasiswaRoute(router) //add this
	routes.MataKuliahRoute(router)

	router.Run("localhost:8080")
}
