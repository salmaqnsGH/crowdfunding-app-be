package main

import (
	"fmt"
	"log"

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

	userInput := user.RegisterUserInput{}
	userInput.Name = "name"
	userInput.Email = "name@email.com"
	userInput.Occupation = "occupation"
	userInput.Password = "password"

	userService.RegisterUser(userInput)

	// input dari user
	// handler : mapping input dari user ke struct input
	// service : mapping dari struct input ke struct user
	// repository : save struct user ke db
	// db
}
