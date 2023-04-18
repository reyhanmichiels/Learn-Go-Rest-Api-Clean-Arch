package main

import (
	"pustaka-api/Database/sql"
	"pustaka-api/Database/sql/seeder"
	"pustaka-api/handler/rest"
	"pustaka-api/repository"
	"pustaka-api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	sql.ConnectDatabase()
	sql.Migrate()
	seeder.Seed()

	repository := repository.Init(sql.DB)
	service := service.Init(repository)
	handler := rest.Init(service)
	handler.Run()

	router := gin.Default()

	router.Run()
}
