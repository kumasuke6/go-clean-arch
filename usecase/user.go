package usecase

import (
	"context"
	"go-clean-arch/ctxkey"
	"go-clean-arch/model"
	"go-clean-arch/repository"
	"go-clean-arch/transaction"
)

type userUsecase struct {
	r           repository.UserRepository
	mr          repository.MessageRepository
	transaction transaction.Transaction
}

func NewUserUsecase(r repository.UserRepository, mr repository.MessageRepository, transaction transaction.Transaction) UserUsecase {
	return &userUsecase{r, mr, transaction}
}

func (c *userUsecase) GetByID(ctx context.Context, id string) (*model.User, error) {
	user, err := c.r.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *userUsecase) Create(ctx context.Context, user *model.User) (string, error) {
	row, err := c.r.Create(ctx, user)
	if err != nil {
		return "", err
	}
	return row, nil
}

func (c *userUsecase) Update(ctx context.Context, user *model.User) error {
	err := c.r.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (c *userUsecase) Delete(ctx context.Context, id string) error {
	_, err := c.transaction.DoInTx(ctx, ctxkey.TxKey, func(ctx context.Context) (interface{}, error) {
		err := c.r.Delete(ctx, id)
		if err != nil {
			return nil, err
		}

		err = c.mr.Delete(ctx, id)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		return err
	}

	return nil
}
