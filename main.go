package main

import (
	"alertapp/price-alert/controllers"
	"fmt"
	// "log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	// "go.mongodb.org/mongo-driver/mongo"
)

// var client *mongo.Client

// type Alert struct {
// 	Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
// 	Lastname  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
// }

func main() {
    r := mux.NewRouter()
	homeController := controllers.NewHomeController(mongoConnect())
	r.HandleFunc("/user/create", homeController.CreateUser).Methods("POST")
	r.HandleFunc("/user/login", homeController.Login).Methods("POST")
	r.HandleFunc("/user/home", homeController.Home).Methods("GET")
    r.HandleFunc("/alert/create", homeController.CreateAlert).Methods("POST")
	r.HandleFunc("/alert/delete", homeController.DeleteAlert).Methods("POST")
	fmt.Println("running on port 8000")
	homeController.Api("wss://stream.binancefuture.com/ws/btcusdt@markPrice")
	
    http.ListenAndServe(":8000", r)
	
}

func mongoConnect() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017/parthAlertDB")
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connection established")
	return session
}