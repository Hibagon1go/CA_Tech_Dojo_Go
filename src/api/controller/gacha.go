// main.goから振られた、ガチャに関するリクエストを実行する機能

package controller

import (
	"api/database"
	"api/middleware"
	"api/model"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Times struct {
	Times int `json:"times"`
}

// キャラクターの抽選を行い、ユーザーの所持キャラクター情報に保存
func Do_Gacha(c *gin.Context) {
	db := database.DBConnect()
	is_Auth, user := middleware.Authorization(c) // まず認証を実行
	var times Times
	if is_Auth {
		// c.ShouldBindJSON(&times)で、POSTされたJSONをtimesにキャスト
		if err := c.ShouldBindJSON(&times); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	results := []map[string]string{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < times.Times; i++ { // times回抽選を実行
		picked_character := PickupCharacter() // キャラクターを抽選

		result := map[string]string{"characterID": "", "name": ""} // response用のマップ
		result["characterID"] = picked_character.CharacterID
		result["name"] = picked_character.CharacterName

		var user_character model.UserCharacter
		// 抽選されたキャラクター情報をuserの所持キャラクターテーブルに保存
		user_character.UserID = user.UserID
		user_character.CharacterName = picked_character.CharacterName
		user_character.UserCharacterID = strconv.Itoa(rand.Intn(100000000)) // db内でユニークなID生成したいが、桁数の大きいrandom数で代用
		user_character.CharacterID = picked_character.CharacterID
		user_character.Value = picked_character.Value

		db.Save(&user_character)
		results = append(results, result)
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
	return
}

// キャラクターを抽選
func PickupCharacter() (picked_character model.Character) {
	db := database.DBConnect()
	box := CharacterBox()
	rand.Seed(time.Now().UnixNano())
	characterID := box[rand.Intn(100)]
	db.First(&picked_character, "character_id=?", characterID)
	picked_character.Value = rand.Intn(10000) // 抽選されたキャラクターにランダムで"価値"を付与

	return picked_character
}
