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
	user := user.User{
		Name: "Test simpan",
	}

	userRepository.Save(user)
}
