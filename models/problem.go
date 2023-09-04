package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Problem struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type string             `bson:"type" json:"type"`
	Name string             `bson:"name" json:"name"`
	Cost int64              `bson:"cost" json:"cost"`
	Time int64
}
