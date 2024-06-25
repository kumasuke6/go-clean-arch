package usecase

import (
	"context"
	"go-clean-arch/model"
	"go-clean-arch/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestUsecaseMessageRead(t *testing.T) {
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
			ctrl := gomock.NewController(t)
			mockMessageRepo := repository.NewMockMessageRepository(ctrl)
			mockMessageRepo.EXPECT().Read(gomock.Any(), tt.user_id).Return(tt.wantMessage, tt.wantErr)

			uc := NewMessageUsecase(mockMessageRepo)
			gotMessage, gotErr := uc.Get(context.Background(), tt.user_id)

			assert.NoError(t, gotErr)
			assert.Equal(t, tt.wantMessage, gotMessage)
		})
	}

	fail := map[string]struct {
		user_id     string
		wantMessage *[]model.Message
		wantErr     error
	}{
		"メッセージ取得失敗": {
			user_id:     "2",
			wantMessage: nil,
			wantErr:     assert.AnError,
		},
	}

	for name, tt := range fail {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockMessageRepo := repository.NewMockMessageRepository(ctrl)
			mockMessageRepo.EXPECT().Read(gomock.Any(), tt.user_id).Return(tt.wantMessage, tt.wantErr)

			uc := NewMessageUsecase(mockMessageRepo)
			gotMessage, gotErr := uc.Get(context.Background(), tt.user_id)

			assert.Error(t, gotErr)
			assert.Equal(t, tt.wantMessage, gotMessage)
		})
	}

}
