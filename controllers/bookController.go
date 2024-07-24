package controllers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/uday778/LibraryManagementSystem-with-gorillaMux/database"

	"github.com/uday778/LibraryManagementSystem-with-gorillaMux/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookIdstr := params["bookId"]

	var RequiredBook models.Book

	id, _ := primitive.ObjectIDFromHex(bookIdstr)
	filter := bson.M{"_id": id}
	err := database.BookCollection.FindOne(context.Background(), filter).Decode(&RequiredBook)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book Doesn't exist"))
		return
	}

	marsheledBook, err := json.Marshal(RequiredBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while marshalling data"))
		w.Write(marsheledBook)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(marsheledBook)

}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	coursor, err := database.BookCollection.Find(context.Background(), bson.M{})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No Books Available"))
		return
	}
	var ListOfBooks models.Books
	for coursor.Next(context.Background()) {
		var book models.Book
		err := coursor.Decode(&book)
		if err != err {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ListOfBooks.AddBookToList(book)
	}
	encodeData, err := json.Marshal(ListOfBooks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while Marshalling the Data"))
		return
	}
	w.WriteHeader(http.StatusFound)
	w.Header().Set("content-type", "application/json")
	w.Write(encodeData)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookstr := params["bookId"]

	id, _ := primitive.ObjectIDFromHex(bookstr)
	filter := bson.M{"_id": id}
	_, err := database.BookCollection.DeleteOne(context.Background(), filter)

	// Bookexists:= database.BookCollection.FindOneAndDelete(context.Background(),filter)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book Doesn't exist"))
		return
	}

	w.Write([]byte("Successfully deleted Book"))

}
func AddBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book

	body, err := io.ReadAll(r.Body)
	if err != nil {

		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &newBook)
	if err != nil {
		http.Error(w, "Error while unmarsheling data", http.StatusInternalServerError)
		return
	}
	existingBook := database.BookCollection.FindOne(context.Background(), bson.M{"book_name": newBook.Book_name})

	if existingBook.Err() == nil {
		w.Write([]byte("User already exists !"))
		return
	}
	newBookToAddInDB := models.Book{
		Id:         primitive.NewObjectID(),
		Book_name:  newBook.Book_name,
		Author:     newBook.Author,
		Created_at: time.Now(),
	}

	_, err = database.BookCollection.InsertOne(context.Background(), newBookToAddInDB)

	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Book added To DB "))
}
