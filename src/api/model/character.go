// キャラクター周りのdbの内容定義

package model

// キャラクターテーブルを定義
type Character struct {
	CharacterName string `json:"characterName"`
	CharacterID   string `json:"characterID"`
	Value         int    `json:"value"`
	Probability   int    `json:"probability"`
}

// ユーザーの所持するキャラクターテーブルを定義
type UserCharacter struct {
	UserID          string
	UserCharacterID string `gorm:"primary_key" json:"userCharacterID"`
	CharacterName   string `json:"characterName"`
	CharacterID     string `json:"characterID"`
	Value           int    `json:"value"`
}
