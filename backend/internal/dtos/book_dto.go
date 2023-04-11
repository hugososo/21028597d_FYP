package dtos

import (
	"comp4913-backend/internal/models"
	"fmt"
)

type GetBookReqDTO struct {
	Address string `json:"address" query:"address" param:"address"`
}

type UploadBookReqDTO struct {
	Owner        string `json:"owner" form:"owner"`
	BookName     string `json:"bookName" form:"bookName"`
	Symbol       string `json:"symbol" form:"symbol"`
	Description  string `json:"description" form:"description"`
	Author       string `json:"author" form:"author"`
	AuthorCity   string `json:"authorCity" form:"authorCity"`
	BookGenre    string `json:"bookGenre" form:"bookGenre"`
	BookLanguage string `json:"bookLanguage" form:"bookLanguage"`
	Edition      string `json:"edition" form:"edition"`
	Keywords     string `json:"keywords" form:"keywords"`
}

func (obj *UploadBookReqDTO) MapToBookMetadata(cid string, coverFile string, bookFile string) *models.BookMetadata {
	metadata := &models.BookMetadata{
		Name:       obj.BookName,
		Image:      fmt.Sprintf("ipfs://%s/%s", cid, coverFile),
		SourceFile: fmt.Sprintf("ipfs://%s/%s", cid, bookFile),
	}

	attributes := []*models.AttributeData{}
	attributes = append(attributes, &models.AttributeData{
		TraitType: models.BookDescriptionKey,
		Value:     obj.Description,
	})
	attributes = append(attributes, &models.AttributeData{
		TraitType: models.BookAuthorKey,
		Value:     obj.Author,
	})
	attributes = append(attributes, &models.AttributeData{
		TraitType: models.BookAuthorCityKey,
		Value:     obj.AuthorCity,
	})
	attributes = append(attributes, &models.AttributeData{
		TraitType: models.BookGenreKey,
		Value:     obj.BookGenre,
	})
	attributes = append(attributes, &models.AttributeData{
		TraitType: models.BookLanguageKey,
		Value:     obj.BookLanguage,
	})
	attributes = append(attributes, &models.AttributeData{
		TraitType: models.BookEditionKey,
		Value:     obj.Edition,
	})
	attributes = append(attributes, &models.AttributeData{
		TraitType: models.BookKeywordsKey,
		Value:     obj.Keywords,
	})
	attributes = append(attributes, &models.AttributeData{
		TraitType: models.BookSymbol,
		Value:     obj.Symbol,
	})
	metadata.Attribute = attributes
	return metadata
}
