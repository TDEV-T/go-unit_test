package main

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=admin password=1234 dbname=golang_test port=5432 sslmode=disable TimeZone=Asia/Bangkok"

type User struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"unique"`
	Age      int
}

func InitializeDB() *gorm.DB {

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&User{})
	return db
}

func AddUser(db *gorm.DB, fullname, email string, age int) error {
	user := User{Fullname: fullname, Email: email, Age: age}

	var count int64
	db.Model(&User{}).Where("email = ?", email).Count(&count)

	if count > 0 {
		return errors.New("email already exists")
	}

	result := db.Create(&user)
	return result.Error
}

func main() {
	db := InitializeDB()
	AddUser(db, "John Doe", "jane.doe@example.com", 30)
}
