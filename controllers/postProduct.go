package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Parthiba-Hazra/products-api/models"
	"github.com/Parthiba-Hazra/products-api/responces"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func PostProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var payload models.Product
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&payload); err != nil {
			var res = responces.NewHTTPResponce(http.StatusBadRequest, err)
			c.JSON(http.StatusBadRequest, res)
			return
		}

		// user the validator to validate required fields
		if validationErr := validate.Struct(&payload); validationErr != nil {
			var res = responces.NewHTTPResponce(http.StatusBadRequest, validationErr)
			c.JSON(http.StatusBadRequest, res)
			return
		}

		var createdAt = time.Now().Format(time.RFC3339)
		newProduct := models.Product{
			GUID:        primitive.NewObjectID(),
			Name:        payload.Name,
			Price:       payload.Price,
			Description: payload.Description,
			CreatedAt:   createdAt,
		}

		result, err := userCollection.InsertOne(ctx, newProduct)
		if err != nil {
			var res = responces.NewHTTPResponce(http.StatusInternalServerError, err)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		var createResponse = responces.NewHTTPResponce(http.StatusCreated, result)

		c.JSON(http.StatusCreated, createResponse)
	}
}
