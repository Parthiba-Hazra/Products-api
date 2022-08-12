package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Parthiba-Hazra/products-api/configs"
	"github.com/Parthiba-Hazra/products-api/models"
	"github.com/Parthiba-Hazra/products-api/responces"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")

func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var products []models.Product
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{})

		if err != nil {
			var res = responces.NewHTTPResponce(http.StatusInternalServerError, err)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		// reading data from dtabase
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleProduct models.Product
			if err = results.Decode(&singleProduct); err != nil {
				var res = responces.NewHTTPResponce(http.StatusInternalServerError, err)
				c.JSON(http.StatusInternalServerError, res)
				return
			}

			products = append(products, singleProduct)
		}

		if len(products) == 0 {
			var res = responces.NewHTTPResponce(http.StatusNotFound, mongo.ErrNoDocuments)
			c.JSON(http.StatusNotFound, res)
			return
		}

		var res = responces.NewHTTPResponce(http.StatusOK, products)

		c.JSON(http.StatusOK, res)
	}
}
