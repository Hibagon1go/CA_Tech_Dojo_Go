// db関連の作業を行う

package database

import (
	"api/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// dbにテーブルが無い時は自動生成し、dbを返す
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Character{})
	return db
}

// dbに接続し、そのdbを返す
func DBConnect() *gorm.DB {
	// 接続先のdbの情報を入力(docker-compose.ymlで定義された)
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "sample"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT) //データベースに接続
	if err != nil {
		panic(err.Error())
	}
	return db
}
