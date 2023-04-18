package sql

import "pustaka-api/entity"

func Migrate() {
	DB.AutoMigrate(entity.Book{})
}
