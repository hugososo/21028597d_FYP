package controllers

import (
	"comp4913-backend/configs"
	"comp4913-backend/internal/dtos"
	"comp4913-backend/internal/services"
	"comp4913-backend/pkg/fileIO"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BookController struct {
	s    *services.BookService
	ipfs *services.IPFSClient
	db   *gorm.DB
}

func NewBookController(s *services.BookService, ipfs *services.IPFSClient, db *gorm.DB) *BookController {
	return &BookController{s, ipfs, db}
}

func (c BookController) GetBooks(ctx echo.Context) error {
	res := c.s.GetBooks()
	return SuccessResponse(ctx, res)
}

func (c BookController) GetBookByAddress(ctx echo.Context) error {
	req := &dtos.GetBookReqDTO{}
	if err := ctx.Bind(req); err != nil {
		return BadResponse(ctx, "")
	}
	res := c.s.GetBookByAddress(req.Address)
	return SuccessResponse(ctx, res)
}

func (c BookController) UploadBook(ctx echo.Context) error {
	req := dtos.UploadBookReqDTO{}
	if err := ctx.Bind(&req); err != nil {
		return BadResponse(ctx, err.Error())
	}
	cover, err := ctx.FormFile("coverImage")
	if err != nil {
		return BadResponse(ctx, err.Error())
	}
	book, err := ctx.FormFile("sourceFile")
	if err != nil {
		return BadResponse(ctx, err.Error())
	}
	cover.Filename = "coverImage"
	book.Filename = "sourceFile"
	// Create folder
	var path = fileIO.CreateTempFolder(req.Symbol)
	defer fileIO.RemoveAll(path)
	coverFile, err := fileIO.AddToFolder(path, cover)
	if err != nil {
		return BadResponse(ctx, err.Error())
	}
	bookFile, err := fileIO.AddToFolder(path, book)
	if err != nil {
		return BadResponse(ctx, err.Error())
	}
	folderCid := c.ipfs.UploadFolder(path)
	metadata := req.MapToBookMetadata(folderCid, coverFile, bookFile)
	fileBytes := fileIO.CreateJsonFormatBytes(path, metadata)
	cid := c.ipfs.UploadItem(fileBytes)
	fmt.Printf("%s%s\n", configs.IPFS_VIEW_URL, cid)
	return SuccessResponse(ctx, cid)
}
