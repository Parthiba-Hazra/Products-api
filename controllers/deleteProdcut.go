package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Parthiba-Hazra/products-api/responces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		guid := c.Param("guid")
		defer cancel()

		productId, _ := primitive.ObjectIDFromHex(guid)

		result, err := userCollection.DeleteOne(ctx, bson.M{"guid": productId})
		if err != nil {
			var res = responces.NewHTTPResponce(http.StatusInternalServerError, err)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		if result.DeletedCount < 1 {
			var res = responces.NewHTTPResponce(http.StatusNotFound, map[string]interface{}{"data": "User with specified ID not found!"})
			c.JSON(http.StatusNotFound, res)
			return
		}

		var deleteRespnse = responces.NewHTTPResponce(http.StatusOK, map[string]interface{}{"data": "User successfully deleted!"})
		c.JSON(http.StatusOK, deleteRespnse)
	}
}
