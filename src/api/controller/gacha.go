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

func Do_Gacha(c *gin.Context) {
	db := database.DBConnect()
	is_Auth, user := middleware.Authorization(c) // まず認証を実行
	var times Times
	if is_Auth {
		// 何回ガチャを引くかの変数であるtimesを取得
		if err := c.ShouldBindJSON(&times); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	user_characters := user.UserCharacters
	results := []map[string]string{}
	for i := 0; i < times.Times; i++ {
		picked_character := PickupCharacter() // characterを抽選

		result := map[string]string{"characterID": "", "name": ""} // response用のマップ
		result["characterID"] = picked_character.CharacterID
		result["name"] = picked_character.CharacterName

		var user_character model.UserCharacter
		// 抽選されたcharacter情報をuserの所持characterテーブルに保存
		rand.Seed(time.Now().UnixNano())
		user_character.UserCharacterID = strconv.Itoa(rand.Intn(100000000)) // db内でユニークなID生成したいが、桁数の大きいrandom数で代用
		user_character.CharacterID = picked_character.CharacterID
		user_character.Value = picked_character.Value

		user_characters = append(user_characters, user_character)
		results = append(results, result)
	}

	db.Model(&user).Update("user_characters", user_characters)
	c.JSON(http.StatusOK, gin.H{"results": results})
	return
}

// characterを抽選
func PickupCharacter() (picked_character model.Character) {
	db := database.DBConnect()
	box := CharacterBox()
	rand.Seed(time.Now().UnixNano())
	characterID := box[rand.Intn(100)]
	db.First(&picked_character, "character_id=?", characterID)
	picked_character.Value = rand.Intn(10000) // 抽選されたcharacterにランダムで"価値"を付与

	return picked_character
}
