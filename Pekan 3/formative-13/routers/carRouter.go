package routers

import (
	"formative-13/controllers"
	"formative-13/config"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	db := config.ConnectDB()
	controllers.SetDB(db)

	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:id", controllers.UpdateCar)
	router.GET("/cars/:id", controllers.GetCar)
	router.DELETE("/cars/:id", controllers.DeleteCar)

	return router
}