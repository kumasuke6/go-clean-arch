package repository

import (
	"context"
	"errors"
	"go-clean-arch/ctxkey"
	"go-clean-arch/infra"
	"go-clean-arch/model"
	"go-clean-arch/query"
	"go-clean-arch/transaction"
	"log"
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRepositoryUserRead(t *testing.T) {
	db := infra.Connect()
	defer db.Close()
	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("../testdata/fixtures/repository/user/read"),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		log.Fatal(err)
	}

	success := map[string]struct {
		id       string
		wantUser *model.User
		wantErr  error
	}{
		"ユーザー取得成功": {
			id: "1",
			wantUser: &model.User{
				ID:        1,
				Name:      "test",
				Age:       20,
				Email:     "test@test.com",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	fail := map[string]struct {
		id       string
		wantUser *model.User
		wantErr  error
	}{
		"ユーザー取得失敗": {
			id:       "2",
			wantUser: nil,
			wantErr:  errors.New("failed to select user: ERROR: syntax error at or near \"fail\" (SQLSTATE 42601)"),
		},
	}

	for name, tt := range success {
		if err := fixtures.Load(); err != nil {
			log.Fatal(err)
		}

		t.Run(name, func(t *testing.T) {
			q := query.NewUserQuery()
			ctx := context.Background()
			repo := NewUserRepository(db, q)
			gotUser, gotErr := repo.Read(ctx, tt.id)

			assert.NoError(t, gotErr)
			assert.Equal(t, tt.wantUser, gotUser)
		})
	}

	for name, tt := range fail {
		if err := fixtures.Load(); err != nil {
			log.Fatal(err)
		}

		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockQuery := query.NewMockUserQuery(ctrl)
			defer ctrl.Finish()
			mockQuery.EXPECT().Read().Return("fail")

			ctx := context.Background()
			repo := NewUserRepository(db, mockQuery)
			gotUser, gotErr := repo.Read(ctx, tt.id)

			assert.EqualError(t, gotErr, tt.wantErr.Error())
			assert.Equal(t, tt.wantUser, gotUser)
		})
	}
}

func TestRepositoryCreate(t *testing.T) {
	db := infra.Connect()
	defer db.Close()

	success := map[string]struct {
		user    *model.User
		wantID  string
		wantErr error
	}{
		"ユーザー作成成功": {
			user: &model.User{
				Name:  "test",
				Age:   20,
				Email: "test@test.com",
			},
			wantID: "1",
		},
	}

	fail := map[string]struct {
		user    *model.User
		wantID  string
		wantErr error
	}{
		"ユーザー作成失敗": {
			user: &model.User{
				Name:  "test",
				Age:   20,
				Email: "test@test.com",
			},
			wantID:  "",
			wantErr: errors.New("failed to create user: ERROR: syntax error at or near \"fail\" (SQLSTATE 42601)"),
		},
	}

	for name, tt := range success {
		t.Run(name, func(t *testing.T) {
			q := query.NewUserQuery()
			ctx := context.Background()
			repo := NewUserRepository(db, q)
			gotID, gotErr := repo.Create(ctx, tt.user)

			assert.NoError(t, gotErr)
			assert.Equal(t, tt.wantID, gotID)
		})
	}

	for name, tt := range fail {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockQuery := query.NewMockUserQuery(ctrl)
			defer ctrl.Finish()
			mockQuery.EXPECT().Create().Return("fail")

			ctx := context.Background()
			repo := NewUserRepository(db, mockQuery)
			gotID, gotErr := repo.Create(ctx, tt.user)

			assert.EqualError(t, gotErr, tt.wantErr.Error())
			assert.Equal(t, tt.wantID, gotID)
		})
	}
}

func TestRepositoryDelete(t *testing.T) {
	db := infra.Connect()
	defer db.Close()
	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("../testdata/fixtures/repository/delete"),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		log.Fatal(err)
	}

	success := map[string]struct {
		id      string
		wantID  string
		wantErr error
	}{
		"ユーザー削除成功": {
			id:     "1",
			wantID: "1",
		},
	}

	fail := map[string]struct {
		id      string
		wantID  string
		wantErr error
	}{
		"ユーザー削除失敗(該当ユーザーなし)": {
			id:      "2",
			wantID:  "",
			wantErr: errors.New("該当ユーザーなし: 2"),
		},
	}

	mockFail := map[string]struct {
		id      string
		wantID  string
		wantErr error
	}{
		"ユーザー削除失敗": {
			id:      "2",
			wantID:  "",
			wantErr: errors.New("failed to delete user: ERROR: syntax error at or near \"fail\" (SQLSTATE 42601)"),
		},
	}

	for name, tt := range success {
		if err := fixtures.Load(); err != nil {
			log.Fatal(err)
		}

		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			q := query.NewUserQuery()
			transaction := transaction.NewTransaction(db)
			repo := NewUserRepository(db, q)

			_, gotErr := transaction.DoInTx(ctx, ctxkey.TxKey, func(ctx context.Context) (interface{}, error) {
				gotErr := repo.Delete(ctx, tt.id)
				return nil, gotErr
			})

			assert.NoError(t, gotErr)

			gotUser, _ := repo.Read(ctx, tt.id)
			assert.Nil(t, gotUser)
		})
	}

	for name, tt := range fail {
		if err := fixtures.Load(); err != nil {
			log.Fatal(err)
		}

		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			q := query.NewUserQuery()
			transaction := transaction.NewTransaction(db)
			repo := NewUserRepository(db, q)

			_, gotErr := transaction.DoInTx(ctx, ctxkey.TxKey, func(ctx context.Context) (interface{}, error) {
				gotErr := repo.Delete(ctx, tt.id)
				return nil, gotErr
			})

			assert.EqualError(t, gotErr, tt.wantErr.Error())
		})
	}

	for name, tt := range mockFail {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockQuery := query.NewMockUserQuery(ctrl)
			defer ctrl.Finish()
			mockQuery.EXPECT().Delete().Return("fail")

			ctx := context.Background()
			transaction := transaction.NewTransaction(db)

			repo := NewUserRepository(db, mockQuery)
			_, gotErr := transaction.DoInTx(ctx, ctxkey.TxKey, func(ctx context.Context) (interface{}, error) {
				gotErr := repo.Delete(ctx, tt.id)
				return nil, gotErr
			})

			assert.EqualError(t, gotErr, tt.wantErr.Error())
		})
	}
}
