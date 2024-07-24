package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	
	"github.com/uday778/LibraryManagementSystem-with-gorillaMux/database"

	"github.com/uday778/LibraryManagementSystem-with-gorillaMux/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func OrderBook(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")
	vars := mux.Vars(r)
	bookIdstr, errbool := vars["bookId"]
	if !errbool {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Book id not provided in url"))
		return
	}
	id, _ := primitive.ObjectIDFromHex(bookIdstr)
	filter := bson.M{"_id": id}
	BookExists := database.BookCollection.FindOne(context.Background(), filter)
	if BookExists.Err() != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book with this id doesn't exist"))
		return
	}

	order_book := models.Order{
		Id:         primitive.NewObjectID(),
		User_id:    username,
		Book_id:    bookIdstr,
		Created_at: time.Now(),
	}

	_, err := database.OrderCollection.InsertOne(context.Background(), order_book)
	if err != nil {
		panic(err)
	}

	w.Write([]byte("Successfully ordered book!"))



}

func Listallorderedbooks(w http.ResponseWriter, r *http.Request)  {
	username := r.Header.Get("username")
	

	coursor, err := database.OrderCollection.Find(context.Background(), bson.M{"user_id": username})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No Ordered Books Available"))
		return
	}
	var ListOfOrders models.Orders
	for coursor.Next(context.Background()) {
		var order models.Order
		err := coursor.Decode(&order)
		if err != err {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ListOfOrders.AddOrderToList(order)
	}
	encodeData, err := json.Marshal(ListOfOrders)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while Marshalling the Data"))
		return
	}
	w.WriteHeader(http.StatusFound)
	w.Header().Set("content-type", "application/json")
	w.Write(encodeData)
	
}