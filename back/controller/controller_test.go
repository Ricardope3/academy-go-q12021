package controller

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ricardope3/academy-go-q12021/back/controller/mocks"
	"github.com/ricardope3/academy-go-q12021/back/models"
	"github.com/stretchr/testify/assert"
)

//mockgen -source=controller/controller.go -destination=controller/mocks/controller.go -package=mocks
func Test_Controller(t *testing.T) {

	tests := []struct {
		name                    string
		expectedParams          []models.Todo
		expectedUsecaseResponse int
		expectUsecaseCall       bool
		expectedError           error
		wantError               bool
	}{
		{
			name:                    "OK, save CSV",
			expectedParams:          []models.Todo{},
			expectedUsecaseResponse: http.StatusAccepted,
			expectUsecaseCall:       true,
			wantError:               false,
			expectedError:           nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			u := mocks.NewMockEntity(mockCtrl)

			if tt.expectUsecaseCall {
				u.EXPECT().SaveCSV(tt.expectedParams).Return(tt.expectedUsecaseResponse)
			}

			c := New(u, nil)

			response := c.entity.SaveCSV(tt.expectedParams)
			assert.Equal(t, response, tt.expectedUsecaseResponse)

		})
	}
}
