package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/salmaqnsGH/crowdfunding-app/auth"
	"github.com/salmaqnsGH/crowdfunding-app/handler"
	"github.com/salmaqnsGH/crowdfunding-app/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=root password=secret dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("koneksi database berhasil!")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
	// input dari user
	// handler : mapping input dari user ke struct input
	// service : mapping dari struct input ke struct user
	// repository : save struct user ke db
	// db
}