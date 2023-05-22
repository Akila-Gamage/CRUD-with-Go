package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//struct to represent the application data
type User struct {
	Id 			primitive.ObjectID `json:"id,omitempty"`
	UserName	string			   `json:"username,omitempty"`
	Email		string			   `json:"email,omitempty"`
	Password	string			   `json:"password,omitempty"`
}