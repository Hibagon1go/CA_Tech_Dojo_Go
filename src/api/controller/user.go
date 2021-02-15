// main.goから振られた、ユーザーに関するリクエストを実行する機能

package controller

import (
	"api/database"
	"api/middleware"
	"api/model"
	"crypto/rand"
	"errors"
	"log"
	"net/http"

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

	randomID, err := MakeRandomStr(32)
	if err != nil {
		log.Fatal(err)
	}
	user.RandomID = randomID // 各ユーザーに固有のランダムな文字列をuserの情報に追加

	token := middleware.CreateToken(user)
	user.Token = token

	// db.Createによりdbにuserを登録
	if err := db.Create(&user).Error; err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

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

		tmp := before_user
		before_user.Name = after_user.Name // userのNameを更新する
		db.Delete(&tmp)
		db.Save(&before_user) // 更新後のユーザーをdbにセーブ
	}
}

// ランダムな文字列生成
func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
