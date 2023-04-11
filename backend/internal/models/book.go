package models

import "github.com/lib/pq"

type Book struct {
	CreatedAt    uint64         `json:"createdAt"`                 //Block number
	Address      string         `json:"address" gorm:"primaryKey"` //book nft contract address
	BookName     string         `json:"bookName"`
	Symbol       string         `json:"symbol"`
	Description  string         `json:"description"`
	Author       string         `json:"author"`
	AuthorCity   string         `json:"authorCity"`
	BookGenre    string         `json:"bookGenre"`
	BookLanguage string         `json:"bookLanguage"`
	Edition      string         `json:"edition"`
	Keywords     pq.StringArray `json:"keywords" gorm:"type:varchar(100)[]"`
	CoverImage   string         `json:"coverImage"`
	SourceFile   string         `json:"sourceFile"`

	Owner   string  `json:"owner"`
	Holders []*User `json:"-" gorm:"many2many:user_bought_books"`
}

const (
	BookSymbol         = "Symbol"
	BookDescriptionKey = "Description"
	BookAuthorKey      = "Author"
	BookAuthorCityKey  = "Author City"
	BookGenreKey       = "Genre"
	BookLanguageKey    = "Language"
	BookEditionKey     = "Edition"
	BookKeywordsKey    = "Keywords"
)
