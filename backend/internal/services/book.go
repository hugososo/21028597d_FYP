package services

import (
	"comp4913-backend/internal/models"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type BookService struct {
	db     *gorm.DB
	client *ethclient.Client
}

func NewBookService(db *gorm.DB, client *ethclient.Client) *BookService {
	return &BookService{db, client}
}

func (s BookService) GetBooks() *[]models.Book {
	books := []models.Book{}
	s.db.Find(&books)
	return &books
}

func (s BookService) GetBookByAddress(address string) *models.Book {
	book := models.Book{}
	s.db.Debug().Model(&book).Where("address = ?", address).First(&book)
	return &book
}

func (s BookService) GetBooksAddress() []string {
	booksAddress := []string{}
	s.db.Model(&models.Book{}).Select("address").Find(&booksAddress)
	return booksAddress
}

func (s BookService) GetBookLastBlock() uint64 {
	book := models.Book{}
	s.db.Order("created_at desc").First(&book)
	return book.CreatedAt
}
