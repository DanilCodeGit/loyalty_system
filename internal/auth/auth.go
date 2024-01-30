package auth

import (
	"fmt"
	"github.com/DanilCodeGit/loyalty_system/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

const TokenExp = time.Hour * 4

var SecretKey = []byte("SecretYouShouldHide")

func GenerateJWT(userLogin domain.Users) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = userLogin.Login
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(SecretKey)

	if err != nil {
		log.Fatalf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(w http.ResponseWriter, r *http.Request) (err error) {
	if r.Header["Token"] == nil {
		fmt.Fprintf(w, "can not find token in header")
		return
	}
	token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.WithMessage(err, "there was an error in parsing")
		}
		return SecretKey, nil
	})

	if token == nil {
		fmt.Fprintf(w, "invalid token")
	}
	return nil
}
