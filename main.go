package main

import (
	"fmt"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"simple-project/internal/application"
	"simple-project/internal/domain"
	"simple-project/internal/interface/api/https"
	"simple-project/internal/interface/persistence/database"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3307)/tokoku?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&domain.User{})
}

func main() {
	InitDB()
	fmt.Println("running")
	InitialMigration()

	e := echo.New()

	userRepo := database.NewUserRepository(DB)
	userUsecase := application.NewUserUsecase(*userRepo)
	userHandler := https.NewUserHandler(userUsecase)

	e.POST("/users", userHandler.AddUser)
	e.GET("/users", userHandler.GetAllUsers)

	e.Logger.Fatal(e.Start(":8080"))
}
