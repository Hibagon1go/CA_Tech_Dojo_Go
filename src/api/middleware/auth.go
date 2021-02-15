// 認証のためのミドルウェア

package middleware

import (
	"api/database"
	"api/model"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTによりtoken作成
func CreateToken(user model.User) string {
	secret := "secret"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
		"randomID": user.RandomID,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// jwtの検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // CreateTokenにて指定した文字列を使います
	})
	if err != nil {
		return token, err
	}
	return token, nil
}

// x-tokenをdbと照合し、認証できたらそのuser情報を返す
func Authorization(c *gin.Context) (is_Auth bool, user model.User) {
	is_Auth = false
	db := database.DBConnect()
	x_token := c.GetHeader("x-token") // リクエストヘッダーからx-tokenを取得

	// x-tokenが無いなら認証失敗
	if x_token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "x-token is empty."})
		return is_Auth, user
	}

	/* token, err := VerifyToken(x_token)

	if err != nil {
		log.Fatal(err)
	}
	*/
	db.First(&user, "token=?", x_token) // x-tokenと一致するtokenを持つレコードをuserに格納させる

	// 与えられたx-tokenを持つuserが存在しないなら認証失敗
	if user.Name == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "Authentication failed."})
		return is_Auth, user
	}

	/* claims, ok := x_token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "Authentication failed."})
		return is_Auth, user
	}

	randomID, ok := claims["randomID"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "Authentication failed."})
		return is_Auth, user
	}

	if user.RandomID != randomID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "Authentication failed."})
		return is_Auth, user
	}
	*/

	is_Auth = true
	return is_Auth, user
}
