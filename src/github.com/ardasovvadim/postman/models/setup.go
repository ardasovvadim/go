package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("mssql", "sqlserver://admin:admin@localhost:1433?database=POSTMAN")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Package{})
	database.AutoMigrate(&Product{})

	DB = database
}
