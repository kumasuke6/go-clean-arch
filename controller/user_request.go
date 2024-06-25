package controller

import (
	"go-clean-arch/model"
	"time"
)

type UserRequest struct {
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func toModel(req UserRequest) *model.User {
	now := time.Now()
	return &model.User{
		Name:      req.Name,
		Age:       req.Age,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
