package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginModel struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type ResponseUserModel struct {
	ID       primitive.ObjectID `bson:"_id"`
	FullName string
	Email    string
}
