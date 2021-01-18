package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	KEY                    string = "JWT-ARY-STARK"
	DEFAULT_EXPIRE_SECONDS int    = 600 //默认过期时间（s）
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"json"`
}

// 创建一个token密钥
func CreateToken() {

}

// JWT -- json web token
// HEADER PAYLOAD SIGNATURE
// This struct is the PAYLOAD
type MyCustomClaims struct {
	User
	jwt.StandardClaims
}

//刷新jwt token
func RefreshToken(tokenString string) (string, error) {
	// first get previous token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(KEY), nil
		})
	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		return "", err
	}
	mySigningKey := []byte(KEY)
	expireAt := time.Now().Add(time.Second * time.Duration(DEFAULT_EXPIRE_SECONDS)).Unix()
	newClaims := MyCustomClaims{
		claims.User,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    claims.User.Name,
			IssuedAt:  time.Now().Unix(),
		},
	}
	// generate new token with new claims
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenStr, err := newToken.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("generate new fresh json web token failed !! error :", err)
		return "", err
	}
	return tokenStr, err
}

//验证jtw token
func ValidateToken(tokenString string) (info User, err error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(KEY), nil
		})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		//fmt.Printf("%v %v", claims.User, claims.StandardClaims.ExpiresAt)
		//fmt.Println("token will be expired at ", time.Unix(claims.StandardClaims.ExpiresAt, 0))
		info = claims.User
	} else {
		fmt.Println("validate tokenString failed !!!", err)
	}
	return
}

//获取jwt token
func GenerateToken(info *User, expiredSeconds int) (tokenString string, err error) {
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}
	// Create the Claims
	mySigningKey := []byte(KEY)
	expireAt := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	fmt.Println("token will be expired at ", time.Unix(expireAt, 0))
	// pass parameter to this func or not
	user := *info
	claims := MyCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    user.Name,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("generate json web token failed !! error :", err)
	} else {
		tokenString = tokenStr
	}
	return
}

// return this result to client then all later request should have header "Authorization: Bearer <token> "
func getHeaderTokenValue(tokenString string) string {
	//Authorization: Bearer <token>
	return fmt.Sprintf("Bearer %s", tokenString)
}
