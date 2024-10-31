package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentData struct {
	StudentCollection *mongo.Collection
}

type Student struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string             `bson:"name"`
	DateOfBirth string             `bson:"date_of_birth"`
	ClassID     primitive.ObjectID `bson:"class_id" json:"class_id"`
	Gender      string             `bson:"gender"`
}
