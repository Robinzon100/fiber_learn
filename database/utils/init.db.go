package dbUtils

import (
	"github.com/Robinzon100/fiber/database"
	migration "github.com/Robinzon100/fiber/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=admin password=admin dbname=oxeni port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.DBConn = db
	migration.Migrations(database.DBConn)
}
