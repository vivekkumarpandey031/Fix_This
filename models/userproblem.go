package models

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserProblem struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UID     string             `bson:"uid" json:"uid,omitempty"`
	Type    string             `bson:"type" json:"type,omitempty"`
	Problem string             `bson:"problem" json:"problem"`
	Brand   string             `bson:"brand" json:"brand"`
	Model   string             `bson:"model" json:"model"`
}

func (u *UserProblem) Validate() error {
	if u.Type == "" {
		return errors.New("empty type field")
	}
	if u.Problem == "" {
		return fmt.Errorf("empty problem field")
	}
	return nil
}
