// テーブルの定義

package model

import (
    "github.com/jinzhu/gorm"
)
// テーブルを定義
type User struct {
    gorm.Model
    Name string `gorm:"size:255"`
}