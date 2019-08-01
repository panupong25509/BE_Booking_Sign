package repositories

import (
	"log"
	"strings"

	"github.com/gobuffalo/buffalo"

	"github.com/JewlyTwin/be_booking_sign/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type Token struct {
	UserID uuid.UUID
	Role   string
	jwt.StandardClaims
}

func GetJWT(c buffalo.Context) (interface{}, interface{}) {
	jwtReq := c.Request().Header.Get("Authorization")
	if jwtReq == "" {
		return nil, models.Error{400, "Not have jwt"}
	}
	return jwtReq, nil
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

func EncodeJWT(user models.User, secret string) string {
	tokenJWT := Token{UserID: user.ID, Role: user.Role}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenJWT)
	jwt, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatalln(err)
	}
	return jwt
}
