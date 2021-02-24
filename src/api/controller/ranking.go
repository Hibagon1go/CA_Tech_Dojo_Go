// main.goから振られた、ランキング関連の機能

package controller

import (
	"api/database"
	"api/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 所持キャラクターの価値の和を基準にした、ユーザーランキングを取得
func TotalRanking(c *gin.Context) {
	db := database.DBConnect()
	var userIDs []string
	db.Select("userID").Find(&userIDs)

	var usercharacters []model.UserCharacter
	ranking := []map[string]string{}
	for i := 0; i < len(userIDs); i++ {
		db.Find(&usercharacters, "userID=?", userIDs[i])
		user_total_value := map[string]string{"name": "", "total_value": ""}
		total_value := 0
		for j := 0; j < len(usercharacters); j++ {
			total_value += usercharacters[j].Value
		}
		user_total_value["total_value"] = strconv.Itoa(total_value)
		var user model.User
		db.Find(&user, "userID=?", usercharacters[0].UserID)
		user_total_value["name"] = user.Name
		ranking = append(ranking, user_total_value)
	}

	//  total_valueでソートしてJSONで返す

}

// 所持キャラクターの価値の最大値を基準にした、ユーザーランキングを取得
func MaxRanking(c *gin.Context) {
	db := database.DBConnect()
	var all_characters []model.Character
	db.Find(&all_characters)

}
