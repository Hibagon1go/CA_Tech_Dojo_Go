// dbの内容定義

package model

// Userテーブルを定義
type User struct {
    Name string `json:"name"`
    Token string `json:"token"`
}