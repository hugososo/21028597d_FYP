package models

import (
	"database/sql"
	"time"
)

type User struct {
	CreatedAt      time.Time      `json:"createdAt"`
	UserAddress    string         `json:"userAddress" gorm:"primaryKey"`
	Name           sql.NullString `json:"name"`
	Email          sql.NullString `json:"email,omitempty"`
	Twitter        sql.NullString `json:"twitter,omitempty"`
	Instagram      sql.NullString `json:"instagram,omitempty"`
	PersonalSite   sql.NullString `json:"personalSite,omitempty"`
	Salt           string         `json:"salt. omitempty"`
	PublishedBooks []*Book        `json:"publishedBooks" gorm:"foreignKey:Owner"`
	BoughtBooks    []*Book        `json:"boughtBooks" gorm:"many2many:user_bought_books"`
}
