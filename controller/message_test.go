package controller

import (
	"go-clean-arch/model"
	"go-clean-arch/usecase"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestControllerMessageRead(t *testing.T) {

	success := map[string]struct {
		userId    string
		wantJson  string
		wantModel *[]model.Message
		wantErr   error
	}{
		"メッセージ取得成功": {
			userId: "1",
			wantModel: &[]model.Message{
				{
					ID:      1,
					UserID:  1,
					Message: "test message",
				},
			},
			wantJson: `[{"id":1,"user_id":1,"message":"test message"}]\n`,
		},
	}

	for name, tt := range success {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/message", strings.NewReader(tt.userId))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockMessageUsecase := usecase.NewMockMessageUsecase(ctrl)
			mockMessageUsecase.EXPECT().Get(gomock.Any(), gomock.Any()).Return(tt.wantModel, tt.wantErr)

			mc := NewMessageContoller(mockMessageUsecase)
			gotErr := mc.Get(c)
			// 改行コードや空白をトリムして比較
			trimmedExpected := strings.TrimSpace(strings.ReplaceAll(tt.wantJson, "\\n", "\n"))
			trimmedActual := strings.TrimSpace(rec.Body.String())
			assert.NoError(t, gotErr)
			assert.Equal(t, trimmedExpected, trimmedActual)
		})
	}
}
