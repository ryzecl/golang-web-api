package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(bookRequest BookRequest, ID int) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	//return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
	//return s.repository.FindByID(ID)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, err := bookRequest.Price.Int64()
	if err != nil {
		return Book{}, err
	}

	rating, err := bookRequest.Rating.Int64()
	if err != nil {
		return Book{}, err
	}

	discount, err := bookRequest.Discount.Int64()
	if err != nil {
		return Book{}, err
	}

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
	//return s.repository.Create(book)
}

func (s *service) Update(bookRequest BookRequest, ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	if err != nil {
		return Book{}, err
	}

	price, err := bookRequest.Price.Int64()
	if err != nil {
		return Book{}, err
	}

	rating, err := bookRequest.Rating.Int64()
	if err != nil {
		return Book{}, err
	}

	discount, err := bookRequest.Discount.Int64()
	if err != nil {
		return Book{}, err
	}

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)

	updateBook, err := s.repository.Update(book)
	return updateBook, err
	//return s.repository.Create(book)
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	if err != nil {
		return Book{}, err
	}

	deletedBook, err := s.repository.Delete(book)
	return deletedBook, err
	//return s.repository.Delete(book)
}
