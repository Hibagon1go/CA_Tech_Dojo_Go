// User周りのdbの内容定義

package model

// Userテーブルを定義
type User struct {
	Name           string          `gorm:"primary_key" json:"name"`
	Token          string          `json:"token"`
	RandomID       string          `json:"randomID"`
	UserCharacters []UserCharacter `json:"userCharacters"`
}
