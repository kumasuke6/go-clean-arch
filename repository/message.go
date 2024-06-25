package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-clean-arch/model"
	"go-clean-arch/query"

	"github.com/go-faster/errors"
)

type messageRepository struct {
	db *sql.DB
	q  query.MessageQuery
}

func NewMessageRepository(db *sql.DB, q query.MessageQuery) MessageRepository {
	return &messageRepository{db, q}
}

func (r *messageRepository) Read(ctx context.Context, user_id string) (*[]model.Message, error) {
	q := r.q.Read()
	rows, err := r.db.Query(q, user_id)
	if err != nil {
		return nil, errors.Errorf("クエリ取得失敗: %w", err)
	}
	defer rows.Close()

	message := []model.Message{} // Initialize the message slice
	for rows.Next() {
		var m model.Message
		err := rows.Scan(&m.ID, &m.UserID, &m.Message)
		if err != nil {
			return nil, errors.Errorf("スキャン失敗: %w", err)
		}
		message = append(message, m) // Append elements to the message slice
	}

	return &message, nil
}

func (r *messageRepository) Create(ctx context.Context, message *model.Message) (string, error) {
	var id string
	err := r.db.QueryRowContext(ctx, "INSERT INTO messages (text) VALUES ($1) RETURNING id", message.Message).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *messageRepository) Delete(ctx context.Context, id string) error {
	return fmt.Errorf("transaction not found")
	_, err := r.db.Exec("DELETE FROM messages WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
