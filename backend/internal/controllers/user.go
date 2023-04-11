package controllers

import (
	"comp4913-backend/internal/dtos"
	"comp4913-backend/internal/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	s  *services.UserService
	db *gorm.DB
}

func NewUserController(s *services.UserService, db *gorm.DB) *UserController {
	return &UserController{s, db}
}

func (c UserController) CreateUser(ctx echo.Context) error {
	req := &dtos.UserReqDTO{}
	if err := ctx.Bind(req); err != nil {
		return BadResponse(ctx, err.Error())
	}
	c.s.CreateUser(req.Address)
	return SuccessResponse(ctx, "Created User")
}

func (c UserController) GetUser(ctx echo.Context) error {
	req := &dtos.UserReqDTO{}
	if err := ctx.Bind(req); err != nil {
		return BadResponse(ctx, err.Error())
	}
	res := c.s.GetUser(req.Address)
	return SuccessResponse(ctx, res)
}

func (c UserController) GetUserWithBooks(ctx echo.Context) error {
	req := &dtos.UserReqDTO{}
	if err := ctx.Bind(req); err != nil {
		return BadResponse(ctx, err.Error())
	}
	res := c.s.GetUserWithBooks(req.Address)
	return SuccessResponse(ctx, res)
}

func (c UserController) GetUserWithPublishes(ctx echo.Context) error {
	req := &dtos.UserReqDTO{}
	if err := ctx.Bind(req); err != nil {
		return BadResponse(ctx, err.Error())
	}
	res := c.s.GetUserWithPublishes(req.Address)
	return SuccessResponse(ctx, res)
}
