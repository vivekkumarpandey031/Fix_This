package models

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Laptop struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Model string             `bson:"model" json:"model"`
	Brand string             `bson:"brand" json:"brand"`
}

func (l *Laptop) Validate() error {
	if l.Model == "" {
		return errors.New("empty laptop model field")
	}
	if l.Brand == "" {
		return fmt.Errorf("empty laptop brand field")
	}
	return nil
}
