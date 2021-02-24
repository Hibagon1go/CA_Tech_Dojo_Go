// まずこれが実行される
package main

import (
	"api/controller"
	"api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	database.DBMigrate(database.DBConnect()) // dbをMigrateし、データ格納のためのテーブルを用意

	controller.RegisterCharacter() // Character情報をdbに登録

	userEngine := engine.Group("/user")
	{
		// controllerへリクエストを振る
		userEngine.POST("/create", controller.CreateUser)
		userEngine.GET("/get", controller.GetUser)
		userEngine.PUT("/update", controller.UpdateUser)

	}

	gachaEngine := engine.Group("/gacha")
	{
		// controllerへリクエストを振る
		gachaEngine.POST("/draw", controller.Do_Gacha)

	}

	characterEngine := engine.Group("/character")
	{
		// controllerへリクエストを振る
		characterEngine.GET("/list", controller.Character_List)

	}

	rankingEngine := engine.Group("/ranking")
	{
		// controllerへリクエストを振る
		rankingEngine.GET("/sum", controller.TotalRanking)
		rankingEngine.GET("/max", controller.MaxRanking)
	}

	engine.Run(":8080") // localhost:8080でサーバー走らせる
}
