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
        userEngine.POST("/create", controller.UserCreate) 
        //userEngine.GET("/get", controller.UserGet)
        //userEngine.PUT("/update", controller.UserUpdate)

    }

    engine.Run(":8080")
}


