package usecase

import (
	"context"
	"go-clean-arch/model"
)

type UserUsecase interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (string, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

type MessageUsecase interface {
	Get(ctx context.Context, id string) (*[]model.Message, error)
	Create(ctx context.Context, message *model.Message) (string, error)
	// Update(ctx context.Context, message *model.Message) error
	Delete(ctx context.Context, id string) error
}
