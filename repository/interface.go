package repository

import (
	"context"
	"go-clean-arch/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (string, error)
	Read(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

type MessageRepository interface {
	Create(ctx context.Context, message *model.Message) (string, error)
	Read(ctx context.Context, user_id string) (*[]model.Message, error)
	Delete(ctx context.Context, id string) error
}
