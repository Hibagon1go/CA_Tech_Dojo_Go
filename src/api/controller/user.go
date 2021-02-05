// main.goから振られたリクエストをserviceに割り振り、レスポンスを返す

package controller

import (
	"github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
	"net/http"
	"api/model"
)

// dbにテーブルが無い時は自動生成し、dbを返す
func DBMigrate(db *gorm.DB) *gorm.DB {
    db.AutoMigrate(&model.User{})
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

func UserCreate(c *gin.Context){
	db := DBMigrate(DBConnect()) 
	name := c.Param("name")
	user := model.User{Name: name}
	err := c.Bind(&user)
	if err != nil{
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
		return
	}

	db.NewRecord(&user)
	db.Create(&user)
    c.JSON(200, gin.H{
		"token" : "Yotto0416",
	})
}

