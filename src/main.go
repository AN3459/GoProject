package main

import (
	"example.com/m/v2/src/controller"
	"example.com/m/v2/src/entity"
	"example.com/m/v2/src/repository"
	"example.com/m/v2/src/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic("failed to migrate database")
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	router := gin.Default()

	router.POST("/users", userController.AddUser)

	err = router.Run(":8080")
	if err != nil {
		panic("failed to start server")
	}
}
