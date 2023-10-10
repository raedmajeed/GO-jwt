package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)


type Credentials struct {
	Username string 	`json:"username"`
	Password string		`json:"password"`		
}

var userData = map[string]string {
	"user1": "1234",
	"user2": "1234",
}

type Claims struct {
	Username string
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err!= nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, ok := userData[credentials.Username]
	if ok == false {
		log.Fatal("User Not Found")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	
	claims := &Claims{
		Username: "raed",
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: time.Now().Add(time.Hour * 1 ).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	sign := []byte("signing")
	signToken, err := token.SignedString(sign)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, 
		&http.Cookie{
			Name: "token",
			Value: signToken,
			Expires: time.Now().Add(time.Minute * 5),
		})
 
}

func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	} 
	
	token := cookie.Value
	claims := &Claims{}
	jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte("signing"), nil
		})
	
	fmt.Println(claims.Username)
}

func Refresh(w http.ResponseWriter, r *http.Request) {

}