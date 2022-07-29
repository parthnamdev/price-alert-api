package models

import "gopkg.in/mgo.v2/bson"

type Alert struct {
	Id 			bson.ObjectId 	`json:"id" bson:"_id"`
	Username 	string			`json:"username" bson:"username"`
	Price 		int				`json:"price" bson:"price"`
	Status 		string			`json:"status" bson:"status"`
}