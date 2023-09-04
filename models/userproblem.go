package models

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserProblem struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UID     string             `bson:"uid" json:"uid"`
	Type    string             `bson:"type" json:"type"`
	Problem string             `bson:"problem" json:"problem"`
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
