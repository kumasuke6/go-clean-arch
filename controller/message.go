package controller

import (
	"go-clean-arch/usecase"

	"github.com/labstack/echo/v4"
)

type messageController struct {
	m usecase.MessageUsecase
}

func NewMessageContoller(m usecase.MessageUsecase) MessageContoller {
	return &messageController{m}
}

func (c *messageController) Get(ctx echo.Context) error {
	user_id := ctx.Param("user_id")
	message, err := c.m.Get(ctx.Request().Context(), user_id)
	if err != nil {
		return ctx.JSON(500, err)
	}
	return ctx.JSON(200, message)
}

func (c *messageController) Create(ctx echo.Context) error {
	// row, err := c.m.Create(ctx, message)
	// if err != nil {
	// 	return "", err
	// }
	return nil
}

func (c *messageController) Delete(ctx echo.Context) error {
	// err := c.m.Delete(ctx, id)
	// if err != nil {
	// 	return err
	// }
	return nil
}
