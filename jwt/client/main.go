package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("vivienvivienvivien")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenarateJWT()

	if err != nil {
		fmt.Fprintf(w, "Fail to generate valid token")
		fmt.Println("Fail to generate valid token")
	}

	fmt.Fprintf(w, validToken)
}

func GenarateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	handleRequest()
}
