// character関連の機能

package controller

import (
	"api/database"
	"api/middleware"
	"api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CharacterID、CharacterName、その排出確率のセットを作り、dbに保存
func RegisterCharacter() {
	db := database.DBConnect()

	var character1 model.Character
	character1.CharacterName = "violin"
	character1.CharacterID = "1"
	character1.Probability = 45
	db.Save(&character1)

	var character2 model.Character
	character2.CharacterName = "viola"
	character2.CharacterID = "2"
	character2.Probability = 30
	db.Save(&character2)

	var character3 model.Character
	character3.CharacterName = "cello"
	character3.CharacterID = "3"
	character3.Probability = 20
	db.Save(&character3)

	var character4 model.Character
	character4.CharacterName = "contrabass"
	character4.CharacterID = "4"
	character4.Probability = 5
	db.Save(&character4)

}

// 排出確率に応じた割合でcharacterが入っている箱を用意(抽選に用いる)
func CharacterBox() (box [100]string) {
	db := database.DBConnect()
	var character model.Character
	idx := 0

	for id := 1; id < 5; id++ {
		s := strconv.Itoa(id)
		db.First(&character, "character_id=?", s)
		for i := 0; i < character.Probability; i++ {
			box[idx+i] = character.CharacterID
		}
		idx += character.Probability
	}
	return box
}

func Character_List(c *gin.Context) {
	db := database.DBConnect()
	is_Auth, user := middleware.Authorization(c) // まず認証を実行
	if is_Auth {
		// userに属する所持character情報を全て取り出す
		character_lists := []map[string]string{}
		var user_characters []model.UserCharacter
		db.Find(&user_characters, "user_id=?", user.UserID)
		for i := 0; i < len(user_characters); i++ {
			character_list := map[string]string{"userCharacterID": "", "characterID": "", "name": ""} // response用のマップ
			character_list["userCharacterID"] = user_characters[i].UserCharacterID
			character_list["characterID"] = user_characters[i].CharacterID
			character_list["name"] = user_characters[i].CharacterName
			character_lists = append(character_lists, character_list)
		}

		c.JSON(http.StatusOK, gin.H{"characters": character_lists})
	}
	return
}
