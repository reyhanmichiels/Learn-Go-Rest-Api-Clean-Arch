package seeder

import (
	"pustaka-api/Database/sql"
	"pustaka-api/entity"
)

func SeedBook() {
	b := entity.Book{
		Title: "Buku satu",
		Price: 20000,
	}
	sql.DB.Create(&b)

	b = entity.Book{
		Title: "Buku dua",
		Price: 30000,
	}
	sql.DB.Create(&b)

	b = entity.Book{
		Title: "Buku tiga",
		Price: 40000,
	}
	sql.DB.Create(&b)
}