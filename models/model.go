package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//struct to represent the application data
type User struct {
	Id 			primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name		string			   `json:"name,omitempty"`
	Email		string			   `json:"email,omitempty"`
	Password	string			   `json:"password,omitempty"`
}

//struct to input login data
type Userlogin struct {
	Name 		string 		`json:"name,omitempty"`
	Password 	string		`json:"password,omitempty"`
}