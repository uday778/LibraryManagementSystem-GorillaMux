package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"io"
	"net/http"

	"github.com/uday778/LibraryManagementSystem-with-gorillaMux/database"

	"github.com/uday778/LibraryManagementSystem-with-gorillaMux/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body,err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w,"Error while Reading Request body" , http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body,&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error while unmarshalling"))
	}
	
	var user_in_db models.User
	err = database.UserCollection.FindOne(context.Background(),bson.M{"username":user.UserName}).Decode(&user_in_db)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user was not found"))
		return 
	}


	//password 
	if user_in_db.Password == user.Password{
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("Valid user"))
	}else{
		w.Write([]byte("Wrong Password"))
	}
	

}
func SignUp(w http.ResponseWriter, r *http.Request) {

	var newUser models.User

	body,err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w,"Error reading request body",http.StatusBadRequest)
		return

	}
	err= json.Unmarshal(body,&newUser)
	if err!= nil{
		http.Error(w,"Error while Unmarshalling data ",http.StatusInternalServerError)
		return
	}

	existingUser := database.UserCollection.FindOne(context.Background(),bson.M{"username":newUser.UserName})

	if existingUser.Err()== nil{
		w.Write([]byte("User already exists"))
		return

	}
	newUserToInsertInDb := models.User{
		Id: primitive.NewObjectID(),
		UserName: newUser.UserName,
		Password: newUser.Password,
		UserType: newUser.UserType,
		Created_at: newUser.Created_at,
	}

	result ,err := database.UserCollection.InsertOne(context.Background(),newUserToInsertInDb)
	if err != nil{
		panic(err)
	}
	fmt.Println("Inserted ID:", result.InsertedID)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte (fmt.Sprintf("New user created with userId %s and password %s",newUser.UserName,newUser.Password)))



}