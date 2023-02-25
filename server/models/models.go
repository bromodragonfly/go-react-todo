package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Создаем схему модели

type ToDoList struct {
	ID primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"` 
	Task string `json:"id_, omitempty"`
	Status bool `json:"status, omitempty"`
}