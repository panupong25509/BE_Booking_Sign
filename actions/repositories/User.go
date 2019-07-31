package repositories

import (
	"encoding/base64"
	"log"
	"reflect"
	"unsafe"

	"golang.org/x/crypto/bcrypt"

	"github.com/JewlyTwin/be_booking_sign/models"
	// "github.com/fwhezfwhez/jwt"

	// "github.com/dgrijalva/jwt-go"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"
)

func Register(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	user := models.User{}
	if !user.CheckParams(data) {
		return nil, models.Error{400, "กรอกข้อมูลไม่ครบ"}
	}
	username, err := GetUserByUsername(c)
	log.Println(username)
	if err == nil {
		// log.Print(err)
		return nil, models.Error{500, "Username นี้มีผู้ใช้แล้ว"}
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data["password"].(string)), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	_ = user.CreateModel(data, string(hash))
	err = db.Create(&user)
	if err != nil {
		return nil, err
	}
	return resSuccess(nil), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserById(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	id, err := uuid.FromString(data["id"].(string))
	if err != nil {
		return nil, err
	}
	user := models.User{}
	err = db.Find(&user, id)
	if err != nil {
		return nil, models.Error{400, "ไม่มีผู้ใช้นี้ใน database"}
	}
	return user, nil
}

func GetUserByIduuid(c buffalo.Context, id uuid.UUID) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	user := models.User{}
	err = db.Find(&user, id)
	// log.Print(user)
	if err != nil {
		return nil, models.Error{400, "ไม่มีผู้ใช้นี้ใน"}
	}
	return user, nil
}

func GetUserByUsername(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	if data["username"] == nil {
		return nil, models.Error{400, "ไม่มี username"}
	}
	username := data["username"].(string)
	user := models.Users{}
	_ = db.Q().Where("username = (?)", username).All(&user)
	if len(user) == 0 {
		return nil, models.Error{400, "ไม่มี username"}
	}
	return user[0], nil
}

func CheckUsernamePassword(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	log.Print("test")
	data := DynamicPostForm(c)
	log.Print("test2")
	username := data["username"].(string)
	log.Print("test2")
	password := data["password"].(string)
	user := models.Users{}
	if username == "" {
		return nil, models.Error{400, "ไม่มี username"}
	}
	if password == "" {
		return nil, models.Error{400, "ไม่มี password"}
	}
	hashBytes, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return nil, err
	}
	_ = db.Q().Where("username = (?)", username).All(&user)
	if CheckPasswordHash(BytesToString(hashBytes), user[0].Password) {
		var secret = "bookingsign"
		tokenString := createTokenString(user[0].ID, secret)
		return tokenString, nil
	}
	return nil, models.Error{400, "ผิดดดดดด"}
}
func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

type Token struct {
	UserID uuid.UUID
	jwt.StandardClaims
}

func createTokenString(userID uuid.UUID, secret string) string {
	// Embed User information to `token`
	tokenJWT := Token{UserID: userID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenJWT)
	// token -> string. Only server knows this secret (foobar).
	tokenstring, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}

func CheckHash(c buffalo.Context) (interface{}, interface{}) {
	db, err := ConnectDB(c)
	if err != nil {
		return nil, err
	}
	data := DynamicPostForm(c)
	password := data["password"].(string)
	hashBytes, err := base64.StdEncoding.DecodeString(password)
	user := models.Users{}
	if password == "" {
		return nil, models.Error{400, "ไม่มี password"}
	}
	_ = db.Q().Where("username = (?)", "panupong").All(&user)
	if CheckPasswordHash(BytesToString(hashBytes), user[0].Password) {
		return &user[0], nil
	}
	// username := data["username"].(string)
	// password := data["password"].(string)
	// claims := jwt.StandardClaims{
	// 	ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
	// 	Id:        "456",
	// }
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// log.Print(token)
	// tokenString, _ := token.SignedString([]byte("hi"))
	// log.Print(tokenString)
	// token.
	// var secret = "kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk"
	// //获取jwt令牌
	// token := jwt.GetToken()
	// token.AddHeader("typ", "JWT").AddHeader("alg", "HS256")
	// exp, err := time.Parse("2006-01-02 15:04:05", "2018-03-20 10:59:44")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, models.Error{400, "ผิดดดดดด"}
	// }
	// token.AddPayLoad("exp", strconv.FormatInt(exp.Unix(), 10))
	// token.AddPayLoad("userid", "ppppp")
	// jwt, _, err := token.JwtGenerator(secret)
	// fmt.Println("签名是:", jwt)

	// p, h, hs, err := token.Decode(jwt)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, models.Error{400, "ผิดดดดดด"}

	// }
	// fmt.Println("payLoad:", p["userid"])

	// fmt.Println("header:", h)
	// fmt.Println("hs256String:", hs)
	return nil, models.Error{400, "ผิดดดดดด"}

}
