package usecase_test

import (
	"context"
	"errors"
	"go-clean-arch/infra"
	"go-clean-arch/repository"
	"go-clean-arch/transaction"
	"go-clean-arch/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestUsecaseUserDelete(t *testing.T) {
	success := map[string]struct {
		id      string
		wantErr error
	}{
		"ユーザー削除成功": {
			id:      "1",
			wantErr: nil,
		},
	}

	fail := map[string]struct {
		id                        string
		wantErrUserDelete         error
		wantErrMessageDelete      error
		wantErrMessageDeleteTimes int
		wantErr                   error
	}{
		"ユーザー削除失敗（ユーザー失敗）": {
			id:                        "1",
			wantErrUserDelete:         errors.New("ユーザー削除失敗"),
			wantErrMessageDelete:      nil,
			wantErrMessageDeleteTimes: 0,
			wantErr:                   errors.New("ユーザー削除失敗"),
		},
		"ユーザー削除失敗（メッセージ失敗）": {
			id:                        "1",
			wantErrUserDelete:         nil,
			wantErrMessageDeleteTimes: 1,
			wantErrMessageDelete:      errors.New("メッセージ削除失敗"),
			wantErr:                   errors.New("メッセージ削除失敗"),
		},
	}

	for name, tt := range success {
		t.Run(name, func(t *testing.T) {
			db := infra.Connect()
			defer db.Close()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := repository.NewMockUserRepository(ctrl)
			mockMessageRepo := repository.NewMockMessageRepository(ctrl)
			mockTransaction := transaction.NewMockTransaction(ctrl)
			ctx := context.Background()

			mockTransaction.EXPECT().DoInTx(ctx, gomock.Any(), gomock.Any()).DoAndReturn(
				func(ctx context.Context, txKey interface{}, fn func(ctx context.Context) (interface{}, error)) (interface{}, error) {
					return fn(ctx)
				},
			)
			mockUserRepo.EXPECT().Delete(gomock.Any(), tt.id).Return(tt.wantErr).Times(1)
			mockMessageRepo.EXPECT().Delete(gomock.Any(), tt.id).Return(tt.wantErr).Times(1)

			uc := usecase.NewUserUsecase(mockUserRepo, mockMessageRepo, mockTransaction)
			gotErr := uc.Delete(ctx, tt.id)

			assert.NoError(t, gotErr)
		})
	}

	for name, tt := range fail {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := repository.NewMockUserRepository(ctrl)
			mockMessageRepo := repository.NewMockMessageRepository(ctrl)
			mockTransaction := transaction.NewMockTransaction(ctrl)
			ctx := context.Background()

			mockTransaction.EXPECT().DoInTx(ctx, gomock.Any(), gomock.Any()).DoAndReturn(
				func(ctx context.Context, txKey interface{}, fn func(ctx context.Context) (interface{}, error)) (interface{}, error) {
					return fn(ctx)
				},
			)

			mockUserRepo.EXPECT().Delete(gomock.Any(), tt.id).Return(tt.wantErrUserDelete)
			mockMessageRepo.EXPECT().Delete(gomock.Any(), tt.id).Return(tt.wantErrMessageDelete).Times(tt.wantErrMessageDeleteTimes)

			uc := usecase.NewUserUsecase(mockUserRepo, mockMessageRepo, mockTransaction)
			gotErr := uc.Delete(ctx, tt.id)

			if tt.wantErr != nil {
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.NoError(t, gotErr)
			}
		})
	}
}
