package main

import (
	config "example/simple_api/config"
	controllers "example/simple_api/controllers"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("config/app.env")

	fmt.Println("Welcome to Rest API Golang!!")
	// config.ConnectDatabase()

	config.ConnectFirebaseStorage()
	r := gin.Default()
	r.GET("api/comment", controllers.All)
	r.GET("api/comment/:id", controllers.Index)
	r.POST("api/comment", controllers.Create)
	r.PUT("api/comment", controllers.Update)
	r.DELETE("api/comment/:id", controllers.Delete)

	r.POST("api/upload", controllers.Upload)
	// r.GET("api/profile", controllers.All)
	// r.GET("api/profile/:id", controllers.Index)
	// r.POST("api/profile", controllers.Create)
	// r.PUT("api/profile", controllers.Update)
	// r.DELETE("api/profile/:id", controllers.Delete)
	r.Run()
}
