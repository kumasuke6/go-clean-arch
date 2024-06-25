package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-clean-arch/ctxkey"
	"go-clean-arch/model"
	"go-clean-arch/query"
	"strconv"

	"github.com/go-faster/errors"
)

type userRepository struct {
	db *sql.DB
	q  query.UserQuery
}

func NewUserRepository(db *sql.DB, q query.UserQuery) UserRepository {
	return &userRepository{db, q}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (string, error) {
	query := r.q.Create()
	result, err := r.db.ExecContext(ctx, query, user.Name, user.Age, user.Email)
	if err != nil {
		return "", errors.Errorf("failed to create user: %w", err)
	}
	rowsAffect, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	rowsAffectStr := strconv.FormatInt(rowsAffect, 10)
	return rowsAffectStr, nil
}

func (r *userRepository) Read(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	query := r.q.Read()
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, errors.Errorf("failed to select user: %w", err)
	}

	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	query := r.q.Update()
	result, err := r.db.Exec(query, user.Name, user.Age, user.Email, user.ID)
	if err != nil {
		return err
	}
	rowsAffect, err := result.LastInsertId()
	if err != nil {
		return err
	}

	if rowsAffect == 0 {
		return fmt.Errorf("no rows affected: %d", user.ID)
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, userID string) error {
	db, ok := ctx.Value(ctxkey.TxKey).(*sql.Tx)
	if !ok {
		return fmt.Errorf("transaction not found")
	}
	query := r.q.Delete()
	result, err := db.Exec(query, userID)
	if err != nil {
		return errors.Errorf("failed to delete user: %w", err)
	}
	rowsAffect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffect == 0 {
		return errors.Errorf("該当ユーザーなし: %s", userID)
	}

	return nil
}
