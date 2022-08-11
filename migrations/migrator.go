package migrations

import (
	//articleData "http/example/features/articles/data"
	userData "http/example/features/users/data"

	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) {
	db.AutoMigrate(userData.User{})
	//db.AutoMigrate(articleData.Article{})
}
