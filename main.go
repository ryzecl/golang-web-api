package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	db.AutoMigrate(&book.Book{})

	// ?CREATE DATA
	// book := book.Book{}
	// book.Title = "Psychology of Money"
	// book.Description = "Buku pengembangan money:)"
	// book.Price = 99000
	// book.Discount = 5
	// book.Rating = 5

	// errCreateBook := db.Create(&book).Error
	// if errCreateBook != nil {
	// }

	// ?READ DATA
	// var books []book.Book

	// readBooks := db.Debug().Where("rating = ?", 5).Find(&books).Error
	// if readBooks != nil {
	// 	fmt.Println("Failed to fetch data", readBooks)
	// 	return
	// }

	// for _, b := range books {
	// 	fmt.Printf("book object %v", b.Title)
	// 	fmt.Printf("book object %v", b)
	// }

	// UPDATE DATA
	// var book book.Book

	// updateBook := db.Debug().Where("id = ?", 1).Find(&book).Error
	// if updateBook != nil {
	// 	fmt.Println("Failed to fetch data", updateBook)
	// 	return
	// }

	// book.Title = "Islam Ala Nabi"
	// errUpdate := db.Save(&book).Error
	// if errUpdate != nil {
	// 	fmt.Println("Failed to update book", errUpdate)
	// 	return
	// }

	// fmt.Printf("book update successfully %v", book)

	// DELETE DATA
	var book book.Book

	deleteBook := db.Debug().Where("id = ?", 3).First(&book).Error
	if deleteBook != nil {
		fmt.Println(deleteBook)
		return
	}

	errDelete := db.Delete(&book).Error
	if errDelete != nil {
		fmt.Printf("Delete failed %v", errDelete)
		return
	}

	fmt.Print("Book deleted successfully")

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8080")
}
