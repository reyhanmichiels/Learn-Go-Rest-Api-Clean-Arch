package main

import (
	"pustaka-api/Database/sql"
	"pustaka-api/Database/sql/seeder"

	"github.com/gin-gonic/gin"
)

func main() {
	sql.ConnectDatabase()
	sql.Migrate()
	seeder.Seed()
	

	// bookRepository := repository.NewBookRepository(db)
	// bookService := service.NewBookService(bookRepository)
	// bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	// v1 := router.Group("api/v1")

	// v1.GET("/books/", bookHandler.GetBookHandler)
	// v1.GET("/books/:id", bookHandler.ShowBookHandler)
	// v1.POST("/books", bookHandler.StoreBookHandler)
	// v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	// v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	router.Run()
}
