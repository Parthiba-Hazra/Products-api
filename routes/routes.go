package routes

import (
	"github.com/Parthiba-Hazra/products-api/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	router.GET("/products", controllers.GetProducts())
	router.GET("/products/:guid", controllers.GetProduct())
	router.POST("/products", controllers.PostProduct())
	router.DELETE("/products/:guid", controllers.DeleteProduct())
	router.PUT("/products/:guid", controllers.PutProduct())
}
