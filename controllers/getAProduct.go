package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Parthiba-Hazra/products-api/models"
	"github.com/Parthiba-Hazra/products-api/responces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		guid := c.Param("guid")
		var product models.Product
		defer cancel()

		productId, _ := primitive.ObjectIDFromHex(guid)

		err := userCollection.FindOne(ctx, bson.M{"guid": productId}).Decode(&product)
		if err != nil {
			var res = responces.NewHTTPResponce(http.StatusNotFound, err)
			c.JSON(http.StatusNotFound, res)
			return
		}

		var getAproductResponse = responces.NewHTTPResponce(http.StatusOK, product)
		c.JSON(http.StatusOK, getAproductResponse)
	}
}
