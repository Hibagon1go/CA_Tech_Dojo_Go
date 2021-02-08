package main

import (
	"api/controller"
	"api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	database.DBMigrate(database.DBConnect()) // dbをMigrateし、データ格納のためのテーブルを用意

	userEngine := engine.Group("/user")
	{
		// controllerへリクエストを振る
		userEngine.POST("/create", controller.CreateUser)
		userEngine.GET("/get", controller.GetUser)
		userEngine.PUT("/update", controller.UpdateUser)

	}

	engine.Run(":8080") // localhost:8080でサーバー走らせる
}
