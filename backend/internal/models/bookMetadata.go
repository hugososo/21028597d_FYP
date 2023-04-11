package models

import (
	"comp4913-backend/internal/contracts"
	"strings"
)

type BookMetadata struct {
	Name       string           `json:"name"`
	Image      string           `json:"image"`
	SourceFile string           `json:"src"`
	Attribute  []*AttributeData `json:"attributes"`
}

type AttributeData struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

func (obj *BookMetadata) MapToBook(lastBlock uint64, eventData contracts.ERC721FactoryBookCreated) *Book {
	book := &Book{
		CreatedAt:    lastBlock,
		Address:      eventData.TokenContract.Hex(),
		BookName:     obj.Name,
		Description:  obj.Attribute[0].Value,
		Author:       obj.Attribute[1].Value,
		AuthorCity:   obj.Attribute[2].Value,
		BookGenre:    obj.Attribute[3].Value,
		BookLanguage: obj.Attribute[4].Value,
		Edition:      obj.Attribute[5].Value,
		Keywords:     strings.Split(obj.Attribute[6].Value, ","),
		CoverImage:   obj.Image,
		SourceFile:   obj.SourceFile,
		Owner:        eventData.Owner.Hex(),
	}
	if len(obj.Attribute) > 7 {
		book.Symbol = obj.Attribute[7].Value
	}
	return book
}
