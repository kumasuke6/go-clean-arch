package main

import (
	"go-clean-arch/controller"
	"go-clean-arch/infra"
	"go-clean-arch/query"
	"go-clean-arch/repository"
	"go-clean-arch/transaction"
	"go-clean-arch/usecase"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomerValidator struct {
	validator *validator.Validate
}

func (cv *CustomerValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomerValidator{validator: validator.New()}

	db := infra.Connect()
	query := query.NewUserQuery()

	transaction := transaction.NewTransaction(db)
	ur := repository.NewUserRepository(db, query)
	mr := repository.NewMessageRepository(db, query)
	uu := usecase.NewUserUsecase(ur, mr, transaction)
	mu := usecase.NewMessageUsecase(mr)
	uc := controller.NewUserContoller(uu)
	mc := controller.NewMessageContoller(mu)

	e.POST("/users", uc.Create)
	e.GET("/users/:id", uc.Get)
	e.PUT("/users", uc.Update)
	e.DELETE("/users/:id", uc.Delete)

	e.GET("/messages/:user_id", mc.Get)

	e.Start(":8080")
}
