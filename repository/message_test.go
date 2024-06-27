package repository

import (
	"context"
	"errors"
	"go-clean-arch/infra"
	"go-clean-arch/model"
	"go-clean-arch/query"
	"log"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestRepositoryMessageRead(t *testing.T) {
	db := infra.Connect()
	defer db.Close()
	q := query.NewMessageQuery()
	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("../testdata/fixtures/repository/message/read"),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		log.Fatal(err)
	}

	success := map[string]struct {
		user_id     string
		wantMessage *[]model.Message
		wantErr     error
	}{
		"メッセージ取得成功": {
			user_id: "1",
			wantMessage: &[]model.Message{
				{
					ID:      1,
					UserID:  1,
					Message: "test message",
				},
			},
		},
		"メッセージ取得成功（2件）": {
			user_id: "2",
			wantMessage: &[]model.Message{
				{
					ID:      2,
					UserID:  2,
					Message: "test message 1",
				},
				{
					ID:      3,
					UserID:  2,
					Message: "test message 2",
				},
			},
		},
	}

	for name, tt := range success {
		t.Run(name, func(t *testing.T) {
			if err := fixtures.Load(); err != nil {
				t.Fatal(err)
			}

			r := NewMessageRepository(db, q)
			ctx := context.Background()
			gotMessage, gotErr := r.Read(ctx, tt.user_id)
			if gotErr != nil {
				t.Fatalf("wantErr: %v, gotErr: %v", tt.wantErr, gotErr)
			}

			assert.Equal(t, tt.wantMessage, gotMessage)
		})
	}

	fail := map[string]struct {
		user_id     string
		wantMessage *[]model.Message
		wantErr     error
	}{
		"メッセージ取得失敗": {
			user_id:     "1",
			wantMessage: nil,
			wantErr:     errors.New("クエリ取得失敗: ERROR: syntax error at or near \"fail\" (SQLSTATE 42601)"),
		},
	}

	for name, tt := range fail {
		if err := fixtures.Load(); err != nil {
			log.Fatal(err)
		}

		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockQuery := query.NewMockMessageQuery(ctrl)
			defer ctrl.Finish()
			mockQuery.EXPECT().Read().Return("fail")

			ctx := context.Background()
			repo := NewMessageRepository(db, mockQuery)
			gotMessage, gotErr := repo.Read(ctx, tt.user_id)

			assert.EqualError(t, gotErr, tt.wantErr.Error())
			assert.Equal(t, tt.wantMessage, gotMessage)
		})
	}
}
