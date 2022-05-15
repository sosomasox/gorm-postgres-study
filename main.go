package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
		if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	userName := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	dbName   := os.Getenv("DB_NAME")
	host     := os.Getenv("HOST")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo", host, userName, password, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	fmt.Println(product)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	fmt.Println(product)

	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	fmt.Println(product)

	// Delete - delete product
	db.Delete(&product, 1)
	fmt.Println(product)
}
