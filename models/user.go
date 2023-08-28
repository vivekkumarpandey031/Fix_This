package models

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `bson:"name" json:"name"`
	Password string `bson:"password" json:"password"`
	Email string `bson:"email" json:"email"`
}

func (u *User)isValidate() error{
	if u.Name==""{
		return errors.New("empty username field")
	}
	if u.Password==""{
		return errors.New("empty password field")
	}
	if u.Email==""{
		return fmt.Errorf("invalid email field")
	}
	
	return nil
}