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

func PutProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		guid := c.Param("guid")
		var product models.Product
		defer cancel()

		productId, _ := primitive.ObjectIDFromHex(guid)

		if err := c.BindJSON(&product); err != nil {
			var res = responces.NewHTTPResponce(http.StatusBadRequest, err)
			c.JSON(http.StatusBadRequest, res)
			return
		}

		if validaionErr := validate.Struct(&product); validaionErr != nil {
			var res = responces.NewHTTPResponce(http.StatusBadRequest, validaionErr)
			c.JSON(http.StatusBadRequest, res)
			return
		}
		var createdAt = time.Now().Format(time.RFC3339)

		update := bson.M{"name": product.Name, "price": product.Price, "description": product.Description, "createdAt": createdAt}
		result, err := userCollection.UpdateOne(ctx, bson.M{"guid": productId}, bson.M{"$set": update})
		if err != nil {
			var res = responces.NewHTTPResponce(http.StatusInternalServerError, err)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		// update the product details
		var updatedProduct models.Product
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"guid": productId}).Decode(&updatedProduct)
			if err != nil {
				var res = responces.NewHTTPResponce(http.StatusInternalServerError, err)
				c.JSON(http.StatusInternalServerError, res)
				return
			}
		}

		var updateRespnse = responces.NewHTTPResponce(http.StatusOK, map[string]interface{}{"data": updatedProduct})
		c.JSON(http.StatusOK, updateRespnse)
	}
}
