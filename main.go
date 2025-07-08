package main

import (
	"GO_Project/controller"
	"GO_Project/database"
	"GO_Project/repository"
	"GO_Project/routes"
	"GO_Project/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database.Connect()
	repo := repository.NewUserRepository(database.DB)
	svc := service.NewUserService(repo)
	ctrl := controller.NewUserController(svc)
	r := routes.SetupRoutes(ctrl)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8081")
}
