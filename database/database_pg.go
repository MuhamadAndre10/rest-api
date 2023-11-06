package database

import (
	"log"
	"rest_api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	DB *gorm.DB
}

// make variable global for access database instance
var Db DbInstance

func ConnecDb() {

	dsn := "host=localhost user=postgres password='root' dbname=product_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("failed to connect to database. \n", err)
	}

	log.Println("connected")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migration")
	db.AutoMigrate(models.Product{})

	Db = DbInstance{
		DB: db,
	}

}
