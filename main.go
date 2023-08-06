package main

import (
	"github.com/codeazq/go-auth-jwt/controllers"
	"github.com/codeazq/go-auth-jwt/initializers"
	"github.com/codeazq/go-auth-jwt/middleware"
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

	r.GET("/validate", middleware.Authenticate, controllers.Validate)

	r.Run()
}
