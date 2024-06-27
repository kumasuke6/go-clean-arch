package controller

import (
	"testing"
)

func TestCreate(t *testing.T) {
	// success := map[string]struct {
	// 	userJson string
	// 	wantErr  error
	// }{
	// 	"ユーザー作成成功": {
	// 		userJson: `{"name":"test","age":20,"email":"test@test"}`,
	// 	},
	// }

	// for name, tt := range success {
	// 	t.Run(name, func(t *testing.T) {
	// 		ctrl := gomock.NewController(t)
	// 		mockUserRepo := usecase.NewMockUserUsecase(ctrl)
	// 		mockUserRepo.EXPECT().Create(gomock.Any(), tt.user).Return(tt.wantId, tt.wantErr)

	// 		uu := usecase.NewUserUsecase(mockUserRepo, nil, nil)
	// 		gotId, gotErr := uu.Create(context.Background(), tt.user)

	// 		assert.NoError(t, gotErr)
	// 		assert.Equal(t, gotId, tt.wantId)
	// 	})
	// }
}
