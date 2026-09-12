package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
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

	book := Book{
		Title: bookRequest.Title,
		Price: int(price),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
	//return s.repository.Create(book)
}
