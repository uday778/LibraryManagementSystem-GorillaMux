package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
	Book_name string `json:"book_name" bson:"book_name"`
	Author string  `json:"author" bson:"author"`
	Created_at time.Time `json:"created_at" bson:"created_at"`
}

type Books struct{
	ListOfBooks []Book
}

func (B *Books)  AddBookToList(book Book){
	B.ListOfBooks=append(B.ListOfBooks, book)
}