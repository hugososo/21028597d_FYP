package controllers

import (
	"comp4913-backend/internal/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type EncryptController struct {
	s  *services.EncryptService
	db *gorm.DB
}

func NewEncryptController(s *services.EncryptService, db *gorm.DB) *EncryptController {
	return &EncryptController{s, db}
}

func (c EncryptController) EncryptFile(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return BadResponse(ctx, "")
	}

	return SuccessResponse(ctx, file)
}
