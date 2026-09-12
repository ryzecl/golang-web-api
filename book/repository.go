package book

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book

	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *repository) FindByID(ID int) (Book, error) {
	var book Book

	err := r.db.Find(&book, ID).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}
