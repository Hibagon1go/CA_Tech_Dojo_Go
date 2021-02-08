// 認証のためのミドルウェア

package middleware

import (
	"api/database"
	"api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// x-tokenをdbと照合し、認証できたらそのuser情報を返す
func Authorization(c *gin.Context) (is_Auth bool, user model.User) {
	is_Auth = false
	db := database.DBConnect()
	token := c.GetHeader("x-token") // リクエストヘッダーからx-tokenを取得
	// x-tokenが無いなら認証失敗
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "x-token is empty."})
		return is_Auth, user
	}

	db.First(&user, "token=?", token) // x-tokenと一致するtokenを持つレコードをuserに格納させる
	// 与えられたx-tokenを持つuserが存在しないなら認証失敗
	if user.Name == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "Authentication failed."})
		return is_Auth, user
	}
	is_Auth = true
	return is_Auth, user
}
