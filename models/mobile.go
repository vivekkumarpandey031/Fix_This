package models

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mobile struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Model string             `bson:"model" json:"model"`
	Brand string             `bson:"brand" json:"brand"`
}

func (m *Mobile) Validate() error {
	if m.Model == "" {
		return errors.New("empty mobile model field")
	}
	if m.Brand == "" {
		return fmt.Errorf("empty mobile brand field")
	}
	return nil
}
