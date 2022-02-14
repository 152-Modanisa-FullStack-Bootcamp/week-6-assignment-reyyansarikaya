package handler

import (
	"bootcamp/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type IHandler interface {
	WalletHandler(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.IWalletService
}

func (h *Handler) WalletHandler(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimPrefix(r.URL.Path, "/")
	if r.Method == http.MethodPut {
		response := h.service.CreateWallet(username)
		jsonData, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
	if r.Method == http.MethodGet {
		if len(username) != 0 {
			response, _ := h.service.OneUserWallet(username)
			jsonData, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		} else {
			response := h.service.ShowAllWallets()
			jsonData, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		}

	}
	if r.Method == http.MethodPost {
		b, err := ioutil.ReadAll(r.Body)

		c := make(map[string]int)
		err = json.Unmarshal(b, &c)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		balance, ok := c["balance"]
		if ok != true {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_ = h.service.UpdateWallet(username, balance)
		w.WriteHeader(http.StatusOK)
	}
}

func NewHandler(service service.IWalletService) IHandler {
	return &Handler{service: service}
}
