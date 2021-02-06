// main.goから振られたリクエストをserviceに割り振り、レスポンスを返す

package controller

import (
	"github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "net/http"
    "log"
    "api/database"
    "api/middleware"
    "api/model"
)

func CreateUser(c *gin.Context){
    db := database.DBMigrate(database.DBConnect()) 
    var user model.User
    // c.ShouldBindJSON(&user)で、POSTされたJSONをuserにキャスト
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }

    token := CreateToken(user)
    user.Token = token 

    // db.Createによりdbにuserを登録
    if err := db.Create(&user).Error; err != nil {
        log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"token" : token,
	})
	
}

// JWTによりtoken作成
func CreateToken(user model.User)(string) {
    var secret string
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
        "name": user.Name,
        "iss": "Y.H", // JWTの発行者が入る
    })

    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
        log.Fatal(err)
    }

    return tokenString
}

func GetUser(c *gin.Context){
    is_Auth, user := middleware.Authorization(c)
    if is_Auth {
        c.JSON(200, gin.H{
            "name" : user.Name,
        })
    }
}