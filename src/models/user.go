package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Username string             `bson:"username" json:"username"`
	Items    []ItemId           `json:"items"`
}

type UserItem struct {
	UserId string `bson:"userId" json:"userId"`
	ItemId string `bson:"itemId" json:"itemId"`
}

type UserItemID struct {
	Id         primitive.ObjectID   `json:"_id"`
	Username   string               `json:"username"`
	ItemOfUser []primitive.ObjectID `json:"itemofuser"`
}
