package models
import (
	"github.com/jinzhu/gorm"
	"github.com/PrathamHandique/go-bookstore/pkg/config"
)
var db *gorm.DB	
type Book struct {
	gorm.Model
	Name string `json:"Name"`
	Author string `json:"Author"`
	Publication string `json:"Publication"`
}
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
func (b *Book) CreateBook() (*Book, error) {
    err := db.Create(&b).Error
    if err != nil {
        return nil, err
    }
    return b, nil
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}
func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID = ?", Id).Delete(book)
	return book
}
