package main

import(
    "github.com/gin-gonic/gin"
    "api/controller"
)

func main() {
    engine := gin.Default()

    userEngine := engine.Group("/user")
    {
        // controllerへリクエストを振る
        userEngine.POST("/create", controller.CreateUser) 
        userEngine.GET("/get", controller.GetUser)
        //userEngine.PUT("/update", controller.UserUpdate)

    }

    engine.Run(":8080") // localhost:8080でサーバー走らせる
}


