// main.goから振られた、ユーザーに関するリクエストを実行する機能

package controller

import (
	"api/database"
	"api/middleware"
	"api/model"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := database.DBConnect()
	var user model.User                // POSTされたユーザー情報を入れる構造体
	var check_already_exist model.User // POSTされたユーザー情報が既に存在するか確認するための構造体

	// c.ShouldBindJSON(&user)で、POSTされたJSONをuserにキャスト
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// dbに既に同じ名前のuserが存在したら登録不可
	db.First(&check_already_exist, "name=?", user.Name)
	if check_already_exist.Name != "" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"response": "This username is already used."})
		return
	}

	token := CreateToken(user)
	user.Token = token

	// db.Createによりdbにuserを登録
	if err := db.Create(&user).Error; err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

// JWTによりtoken作成
func CreateToken(user model.User) string {
	var secret string
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"iss":  "Y.H", // JWTの発行者が入る
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

func GetUser(c *gin.Context) {
	is_Auth, user := middleware.Authorization(c)
	if is_Auth {
		c.JSON(http.StatusOK, gin.H{
			"name": user.Name,
		})
	}
}

func UpdateUser(c *gin.Context) {
	db := database.DBConnect()
	is_Auth, before_user := middleware.Authorization(c) // まず認証を実行
	var after_user model.User
	// c.ShouldBindJSON(&after_user)で、PUTされたJSONをafter_userにキャスト
	if is_Auth {
		if err := c.ShouldBindJSON(&after_user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var tmp model.User
		db.First(&tmp, "name=?", before_user.Name) // 更新前のユーザー情報をdbから受取りtmpに入れる
		tmp.Name = after_user.Name                 // tmpのNameを更新する
		db.Delete(&before_user)                    // 更新前のユーザーをdbから消去
		db.Save(&tmp)                              // 更新後のユーザーをdbにセーブ
	}
}
