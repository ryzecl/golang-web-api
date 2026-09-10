package book

import "gorm.io/gorm"

type Repository interface {
	findAll() ([]Book, error)
	findByID(ID int) (Book, error)
	create(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

