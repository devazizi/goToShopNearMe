package adapter

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQL_DB struct {
	store *gorm.DB
}

func InitDB(dsn string) SQL_DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("can't connect to database")
	}

	var entities = []any{}

	if err := db.AutoMigrate(entities...); err != nil {
		panic("migration fail")
	}

	return SQL_DB{store: db}
}
