package controllers

import (
	"alertapp/price-alert/models"
	"encoding/json"
	"fmt"
	"net/http"

	// "github.com/gorilla/mux"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var jwtKey = []byte("secret_key")

type HomeController struct {
	session *mgo.Session
}

func NewHomeController(session *mgo.Session) *HomeController {
	return &HomeController{session}
}

func (hc HomeController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Running")
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)

	user.Id = bson.NewObjectId()
	hc.session.DB("mongo-golang").C("users").Insert(user)

	userJson, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(w, "%s\n", userJson)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", `{
		"success": "true",
		"message": "user created",
		"data": {},
	}`)
	// vars := mux.Vars(r)
	// username := vars["username"]
	// email := vars["email"]

	// fmt.Fprintf(w, "You've requested the book: %s on page %s\n", alert_price, status)
}

func (hc HomeController) CreateAlert(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// alert_price := vars["price"]
	// status := "created"
	// fmt.Println("Running")
	alert := models.Alert{}
	json.NewDecoder(r.Body).Decode(&alert)

	alert.Id = bson.NewObjectId()
	hc.session.DB("mongo-golang").C("alerts").Insert(alert)

	userJson, err := json.Marshal(alert)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(w, "%s\n", userJson)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// fmt.Fprintf(w, "%s\n", `{
	// 	"success": "true",
	// 	"message": "user created",
	// 	"data": {},
	// }`)
	json.NewEncoder(w).Encode(alert)
	// fmt.Fprintf(w, "You've requested the book: %s on page %s\n", alert_price, status)
}

func (hc HomeController) DeleteAlert(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["id"]
	// status := "deleted"
	alert := models.Alert{}
	json.NewDecoder(r.Body).Decode(&alert)

	alert.Status = "deleted"
	hc.session.DB("mongo-golang").C("alerts").Update(bson.M{"_id": alert.Id}, alert)

	userJson, err := json.Marshal(alert)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(w, "%s\n", userJson)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// fmt.Fprintf(w, "%s\n", `{
	// 	"success": "true",
	// 	"message": "user created",
	// 	"data": {},
	// }`)
	json.NewEncoder(w).Encode(alert)
	// fmt.Fprintf(w, "You've requested the book: %s on page %s\n", alert_price, status)
}

func (hc HomeController) Login(w http.ResponseWriter, r *http.Request) {
	// var credentials Credentials
	// err := json.NewDecoder(r.Body).Decode(&credentials)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	vars := make(map[string]string)
	json.NewDecoder(r.Body).Decode(&vars)
	username := vars["username"]
	password := vars["password"]
	fmt.Println(vars)
	fmt.Println(password)
	result := models.User{}
	hc.session.DB("mongo-golang").C("users").Find(bson.M{"username": username, "password": password}).One(&result)
	// expectedPassword, ok :=
	fmt.Println(result)

	if result.Id == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 500)

	claims := &models.Claims{
		Username: username,
		Email:    result.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	fmt.Println(tokenString)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", `{
		"success": "true",
		"message": "user created",
		"data": {},
	}`)

}

func (hc HomeController) Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	result := []models.Alert{}
	hc.session.DB("mongo-golang").C("alerts").Find(bson.M{}).All(&result)
	// fmt.Fprintf(w, "%s\n", `{
	// 	"success": "true",
	// 	"message": "data fetched",
	// 	"data": {
	// 		"result":
	// 	},
	// }`)
	json.NewEncoder(w).Encode(result)
}
