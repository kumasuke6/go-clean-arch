package controller

import (
	"go-clean-arch/usecase"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userContoller struct {
	u usecase.UserUsecase
}

func NewUserContoller(u usecase.UserUsecase) UserContoller {
	return &userContoller{u}
}

func (c *userContoller) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	u, err := c.u.GetByID(ctx.Request().Context(), id)
	if err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, u)
}

func (c *userContoller) Create(ctx echo.Context) error {
	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	u := toModel(req)

	id, err := c.u.Create(ctx.Request().Context(), u)
	if err != nil || id == "" {
		log.Print(err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, "success")
}

func (c *userContoller) Update(ctx echo.Context) error {
	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(req); err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	u := toModel(req)
	c.u.Update(ctx.Request().Context(), u)
	return nil
}

func (c *userContoller) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if err := c.u.Delete(ctx.Request().Context(), id); err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return nil
}
