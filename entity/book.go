package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string `json:"title" gorm:"type:varchar(255);unique"`
	Price int    `json:"price" gorm:"type:float"`
}

type BookResponse struct {
	Title string `json:"title"`
	Price int    `json:"price"`
}

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required,numeric"`
}
