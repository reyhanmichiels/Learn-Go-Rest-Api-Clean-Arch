package rest

import (
	"pustaka-api/service"

	"github.com/gin-gonic/gin"
)

type Rest interface {
	Run()
}

type rest struct {
	gin *gin.Engine
	service *service.Service
}

func Init(ser *service.Service) Rest {
	r := &rest{
		gin: gin.Default(),
		service: ser,
	}

	r.Router()
	return r
}

func (r *rest) Router() {
	v1 := r.gin.Group("api/v1")

	v1.GET("/books/", r.GetBookHandler)
	v1.GET("/books/:id", r.ShowBookHandler)
	v1.POST("/books", r.StoreBookHandler)
	v1.PUT("/books/:id", r.UpdateBookHandler)
	v1.DELETE("/books/:id", r.DeleteBookHandler)
}

func(r *rest) Run() {
	r.gin.Run()
}
