package models

import (
"go.mongodb.org/mongo-driver/bson/primitive"
"time"
)

type User struct{
	ID primitive.ObjectID `bson:"_id"`
	First_name *string  `json:"first_name" validate:"required,min=2,max=100"`
	LAst_name     *string `json:"last_name" validate:"required,main=2,max=100"`
	Password      *string  `json:"Password" validate:"required, min=6"`
	Email         *string  `json:"email" validation:"email,required"`
	Phone         *string  `json:"phone" validation:"required"`
	Token         *string  `json:"token"`
	User_type     *string  `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token *string   `json:"refresh_token"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	USer_id       string     `json:"user_id"`
}