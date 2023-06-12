package controllers

import (
	"context"
	"crud-with-go/configs"
	"crud-with-go/models"
	"crud-with-go/responses"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

//function to create user
func CreateUser(c echo.Context) error {			
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	var user models.User
	defer cancel()

	//bind the request body to user
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	newUser := models.User {
		Id: primitive.NewObjectID(),
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})

}

//function to get/read a user
func GetUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	userId := c.Param("userId")
	var user models.User
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(userId)

	err := userCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})	
	}

	return c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": user}})
}

// function to Edit/Update a user
func EditUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	userId := c.Param("userId")
	var user models.User
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(userId)

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	update := bson.M{"name": user.Name, "email": user.Email, "password": user.Password}

	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": update})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	var updatedUser models.User
	if result.MatchedCount == 1 {
		err = userCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&updatedUser)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}
	}

	return c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": updatedUser}})
}

//function to delete a user
func DeleteUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	userId := c.Param("userId")
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(userId)

	result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objectId})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.JSON(http.StatusNotFound, responses.Response{Status: http.StatusNotFound, Message: "not found", Data: &echo.Map{"data": "user not found"}})
	}

	return c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": "User deleted"}})
}

//function to Get all the users
func GetAllUsers(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	var users []models.User
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
					return c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}

		users = append(users, singleUser)
	}
	return c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": users}})
}

//function to login in flutter
func Login(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	var userlogin models.Userlogin
	defer cancel()

	//bind the request body to user
	if err := c.Bind(&userlogin); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	newUser := models.Userlogin {
		Name: userlogin.Name,
		Password: userlogin.Password,
	}

	err := userCollection.FindOne(ctx, newUser).Decode(&userlogin)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": userlogin}})
}