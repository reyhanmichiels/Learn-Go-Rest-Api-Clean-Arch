package sql

import "pustaka-api/entity"

func Migrate() {
	DB.Migrator().DropTable(entity.Book{})
	DB.AutoMigrate(entity.Book{})
}
