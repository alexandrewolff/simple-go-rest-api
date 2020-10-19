package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Car struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Brand string             `json:"brand" bson:"brand,omitempty"`
	Model string             `json:"model" bson:"model,omitempty"`
	Year  string             `json:"year" bson:"year,omitempty"`
}
