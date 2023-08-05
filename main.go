package main

import (
	"github.com/codeazq/go-auth-jwt/controllers"
	"github.com/codeazq/go-auth-jwt/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)

	r.POST("/login", controllers.Login)

	r.Run() // listen and serve on 0.0.0.0:8080
}
