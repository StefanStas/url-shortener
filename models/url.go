package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Url struct {
    Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
    Url       string             `json:"url" bson:"url"`
    Hash      string             `json:"hash" bson:"hash"`
    CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
    UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
