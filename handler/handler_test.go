package handler

import (
	mock_service "bootcamp/mock"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerWallets(t *testing.T) {
	t.Run("Get - All Wallet", func(t *testing.T) {
		service := mock_service.NewMockIWalletService(gomock.NewController(t))
		serviceResult := map[string]int{"fatma": 10, "reyyan": 10}
		service.EXPECT().ShowAllWallets().Return(serviceResult).Times(1)

		handler := NewHandler(service)

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.WalletHandler(w, r)

		response := "{\"fatma\":10,\"reyyan\":10}"

		assert.Equal(t, w.Body.String(), response)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "", w.Header().Get("content-type"))
	})
	t.Run("Put - Create Wallet", func(t *testing.T) {
		service := mock_service.NewMockIWalletService(gomock.NewController(t))
		serviceResult := "Wallet created"

		service.EXPECT().CreateWallet("reyyan").Return(serviceResult).Times(1)
		handler := NewHandler(service)

		r := httptest.NewRequest(http.MethodPut, "/reyyan", nil)
		w := httptest.NewRecorder()
		handler.WalletHandler(w, r)

		response := "\"Wallet created\""
		assert.Equal(t, w.Body.String(), response)
		//assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	})
	t.Run("GetWalletByName", func(t *testing.T) {
		service := mock_service.NewMockIWalletService(gomock.NewController(t))
		serviceResult := 10
		service.EXPECT().OneUserWallet("reyyan").Return(serviceResult, true).Times(1)

		handler := NewHandler(service)

		r := httptest.NewRequest(http.MethodGet, "/reyyan", nil)
		w := httptest.NewRecorder()
		handler.WalletHandler(w, r)

		response := "10"
		assert.Equal(t, w.Body.String(), response)
		assert.Equal(t, w.Result().StatusCode, 200)
		assert.Equal(t, "", w.Header().Get("content-type"))

	})
	t.Run("Post - Update Wallet", func(t *testing.T) {
		service := mock_service.NewMockIWalletService(gomock.NewController(t))

		service.EXPECT().UpdateWallet("reyyan", 100).Return("Transaction successful.").Times(1)

		handler := NewHandler(service)
		body := map[string]int{"balance": 100}
		lastBody, _ := json.Marshal(body)

		r := httptest.NewRequest(http.MethodPost, "/reyyan", nil)
		w := httptest.NewRecorder()

		handler.WalletHandler(w, r)
		assert.Equal(t, w.Body.String(), lastBody)
	})
}
