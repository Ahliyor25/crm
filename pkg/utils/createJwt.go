package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(user uint) (access string, retry string, err error) {
	AtExpires := time.Now().Add(time.Hour * 6).Unix()

	token := jwt.New(jwt.SigningMethodHS512)
	atClaims := token.Claims.(jwt.MapClaims)

	//Creating Access Token
	atClaims["authorized"] = false
	atClaims["user_id"] = user
	atClaims["exp"] = AtExpires

	fmt.Println(atClaims)
	AccessToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", "", err
	}

	//Creating Refresh Token
	rtClaims := token.Claims.(jwt.MapClaims)
	RtExpires := time.Now().Add(time.Hour * (24 * 30 * 12 * 100)).Unix()

	rtClaims["authorized"] = true
	rtClaims["user_id"] = user
	rtClaims["exp"] = RtExpires
	rtClaims["token"] = AccessToken

	RefreshToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return AccessToken, RefreshToken, nil
}
