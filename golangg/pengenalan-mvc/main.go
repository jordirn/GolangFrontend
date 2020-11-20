package main

import (
	"github.com/hairunisa29/pengenalan-mvc/tree/master/app/controller"
	"github.com/gin-gonic/gin"
)



func main(){
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.POST("/api/v1/antrian", controller.AddAntrianHandler)
	router.GET("/api/v1/antrian/status", controller.GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian", controller.UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/id/:idAntrian/delete", controller.DeleteAntrianHandler)
	router.GET("/antrian", controller.PageAntrianHandler)
	router.Run(":8080")
}