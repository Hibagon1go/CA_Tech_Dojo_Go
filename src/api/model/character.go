// Character周りのdbの内容定義

package model

// Characterテーブルを定義
type Character struct {
	CharacterName string `json:"characterName"`
	CharacterID   string `json:"characterID"`
	Value         int    `json:"value"`
	Probability   int    `json:"probability"`
}

// Userの所持するCharacterテーブルを定義
type UserCharacter struct {
	UserID          string
	UserCharacterID string `gorm:"primary_key" json:"userCharacterID"`
	CharacterName   string `json:"characterName"`
	CharacterID     string `json:"characterID"`
	Value           int    `json:"value"`
}
