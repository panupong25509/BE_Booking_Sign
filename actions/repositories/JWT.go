package repositories

import (
	"log"
	"strings"

	"github.com/JewlyTwin/be_booking_sign/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type Token struct {
	UserID uuid.UUID
	jwt.StandardClaims
}

func DecodeJWT(jwtReq string, key string) (jwt.MapClaims, interface{}) {
	jwtStrings := strings.Split(jwtReq, "Bearer ")
	token, err := jwt.Parse(jwtStrings[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, models.Error{500, "Token error."}
	}
	tokens := token.Claims.(jwt.MapClaims)
	return tokens, nil
}

func EncodeJWT(userID uuid.UUID, secret string) string {
	tokenJWT := Token{UserID: userID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenJWT)
	jwt, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatalln(err)
	}
	return jwt
}
