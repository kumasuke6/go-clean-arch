package usecase

import (
	"context"
	"go-clean-arch/model"
	"go-clean-arch/repository"
)

type messageUsecase struct {
	mr repository.MessageRepository
}

func NewMessageUsecase(mr repository.MessageRepository) MessageUsecase {
	return &messageUsecase{mr}
}

func (c *messageUsecase) Get(ctx context.Context, id string) (*[]model.Message, error) {
	message, err := c.mr.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (c *messageUsecase) Create(ctx context.Context, message *model.Message) (string, error) {
	row, err := c.mr.Create(ctx, message)
	if err != nil {
		return "", err
	}
	return row, nil
}

func (c *messageUsecase) Delete(ctx context.Context, id string) error {
	err := c.mr.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
