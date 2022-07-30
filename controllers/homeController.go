package controllers

import (
	"alertapp/price-alert/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// "github.com/gorilla/mux"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
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
	// fmt.Println(claims.Username);
	
	result := []models.Alert{}
	hc.session.DB("mongo-golang").C("alerts").Find(bson.M{"username": claims.Username}).All(&result)
	// fmt.Fprintf(w, "%s\n", `{
	// 	"success": "true",
	// 	"message": "data fetched",
	// 	"data": {
	// 		"result":
	// 	},
	// }`)

	
	

	json.NewEncoder(w).Encode(result)
}

// type Trade struct {
// 	Exchange  string  `json:"exchange"`
// 	Base      string  `json:"base"`
// 	Quote     string  `json:"quote"`
// 	Direction string  `json:"direction"`
// 	Price     float64 `json:"price"`
// 	Volume    int64   `json:"volume"`
// 	Timestamp int64   `json:"timestamp"`
// 	PriceUsd  float64 `json:"priceUsd"`
// }

type Trade struct {
    e  string// Event type
    E int         // Event time
    s string         // Symbol
	p string      // Mark price
    i string //     // Index price
    P string //      // Estimated Settle Price, only useful in the last hour before the settlement starts
    r string //        // Funding rate
    T int //         // Next funding time
  }

func (hc HomeController) Api(url string) {
	// response, err := http.Get(url)
	// cookie, err := r.Cookie("token")
	// if err != nil {
	// 	if err == http.ErrNoCookie {
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// tokenStr := cookie.Value

	// claims := &models.Claims{}

	// tkn, err := jwt.ParseWithClaims(tokenStr, claims,
	// 	func(t *jwt.Token) (interface{}, error) {
	// 		return jwtKey, nil
	// 	})

	// if err != nil {
	// 	if err == jwt.ErrSignatureInvalid {
	// 		// w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}
	// 	// w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// if !tkn.Valid {
	// 	// w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }
	// fmt.Println(claims.Username);
	// result := []models.Alert{}
	// hc.session.DB("mongo-golang").C("alerts").Find(bson.M{"username": claims.Username}).All(&result)

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		panic(err)
	}
	input := make(chan Trade)                   

	go func() {                                 
		// read from the websocket
		for {
		_, message, err := c.ReadMessage()   
		if err != nil {
			break
		}
		// unmarshal the message
		var trade Trade
		json.Unmarshal(message, &trade)      
		// send the trade to the channel
		// fmt.Println(trade)
		
		input <- trade         
		}
		close(input)   
		                     
	}()
	defer c.Close()
	
	for trade := range input {
		json.Marshal(trade)
		result := []models.Alert{}
		hc.session.DB("mongo-golang").C("alerts").Find(bson.M{"status": "created"}).All(&result)

		for _, element := range result {
			i, _ := strconv.ParseFloat(trade.P, 32)
			// fmt.Println(trade.P, int(i))
			fmt.Println("price not reached, current price", i, " alert set by user : ", element.Username, " at : ", element.Price)
			if element.Price == int(i){
				us := models.User{}
				hc.session.DB("mongo-golang").C("users").Find(bson.M{"username": element.Username}).One(&us)
				fmt.Println("price reached", element.Price, " alert set by user : ", element.Username, " , email : ", us.Email)
				
				// ms := mailSender.NewMailNotify(host, port , username, password, "Price reach notification", "parthnamdevpm12345@gmail.com", "soham2112@gmail.com")
				// ms.Send("Bitcoin price has reached to your set status")
				
			}
		}
		
	}
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return 
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}


func wsEndpoint(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    // upgrade this connection to a WebSocket
    // connection
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
    }

    log.Println("Client Connected")
    err = ws.WriteMessage(1, []byte("Hi Client!"))
    if err != nil {
        log.Println(err)
    }
    // listen indefinitely for new messages coming
    // through on our WebSocket connection
    reader(ws)
}
func reader(conn *websocket.Conn) {
    for {
    // read in a message
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
    // print out that message for clarity
        fmt.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }

    }
}