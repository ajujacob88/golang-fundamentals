package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=ajujacob dbname=new_test_db2 port=5432 sslmode=disable TimeZone=Asia/kolkata"

	gorm.Open(postgres.Open(dsn))

}
