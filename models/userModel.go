package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"user_id"`
	UserName   string             `json:"username" bson:"username"`
	Password   string             `bson:"password" json:"password"`
	UserType   string  `json:"userType" bson:"userType"`           
	Created_at time.Time       `json:"created_at,omitempty" bson:"created_at"`
}
