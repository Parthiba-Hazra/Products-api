package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	GUID        primitive.ObjectID `json:"guid,omitempty"`
	Name        string             `json:"name" validate:"required"`
	Price       float64            `json:"price" validate:"required"`
	Description string             `json:"description" validate:"required,min=2,max=200"`
	CreatedAt   string             `json:"createdAt,omitempty"`
}
