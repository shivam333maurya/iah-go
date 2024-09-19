package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Test struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
