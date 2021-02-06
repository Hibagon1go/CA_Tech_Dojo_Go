package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"api/database"
	"api/model"
)

// x-tokenをdbと照合し、認証できたらそのuser情報を返す
func Authorization(c *gin.Context) (is_Auth bool, user model.User){
	is_Auth = false
	db := database.DBConnect()
    token := c.GetHeader("x-token") // リクエストヘッダーからx-tokenを取得
    if token == "" {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "x-token is empty."})
        return is_Auth, user
	}
	
	db.First(&user, "token=?", token) // x-tokenと一致するtokenを持つレコードをuserに格納させる
	if user.Name == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "No such user."})
		return is_Auth, user	
	}
	is_Auth = true
	return is_Auth, user
}