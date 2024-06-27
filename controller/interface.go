package controller

import "github.com/labstack/echo/v4"

type UserContoller interface {
	Get(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type MessageContoller interface {
	Get(ctx echo.Context) error
	Create(ctx echo.Context) error
	// Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
