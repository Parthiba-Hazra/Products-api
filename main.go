package main

import (
	"github.com/Parthiba-Hazra/products-api/configs"
	"github.com/Parthiba-Hazra/products-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	routes.UserRoute(router)

	router.Run("localhost:3000")
}
