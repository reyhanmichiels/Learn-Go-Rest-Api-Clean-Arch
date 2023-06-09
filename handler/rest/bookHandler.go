package rest

import (
	"fmt"
	"log"
	"pustaka-api/entity"
	"pustaka-api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *bookHandler {
	return &bookHandler{bookService}
}

func (h *rest) GetBookHandler(c *gin.Context) {
	book, err := h.service.Book.GetBook()
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "succesed",
		"message": "succesfully get all book data",
		"data":    book,
	})
}

func (h *rest) ShowBookHandler(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	book, err := h.service.Book.ShowBook(uint(ID))
	if err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "succesed",
		"message": fmt.Sprintf("succesfully show book with id %d", ID),
		"data":    book,
	})
}

func (h *rest) StoreBookHandler(c *gin.Context) {
	var book entity.BookInput

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(409, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	b, err := h.service.Book.StoreBook(book)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"status":  "succesed",
		"message": "succesfully create new book data",
		"data":    b,
	})
}

func (h *rest) UpdateBookHandler(c *gin.Context) {
	var bookInput entity.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		c.JSON(409, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	b, err := h.service.Book.UpdateBook(uint(ID), bookInput)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "succesed",
		"message": "succesfully update book data",
		"data":    b,
	})
}

func (h *rest) DeleteBookHandler(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	_, err = h.service.Book.ShowBook(uint(ID))
	if err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err = h.service.Book.DeleteBook(uint(ID))
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "succesed",
		"message": fmt.Sprintf("succesfully delete book with id %d", ID),
		"data":    nil,
	})
}
