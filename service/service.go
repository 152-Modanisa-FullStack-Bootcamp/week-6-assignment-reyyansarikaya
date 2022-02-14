package service

import (
	"bootcamp/data"
)

type IWalletService interface {
	CreateWallet(string) string
	OneUserWallet(string) (int, bool)
	ShowAllWallets() map[string]int
	UpdateWallet(string, int) string
}

type WalletService struct {
	data data.IData
}

func (s *WalletService) CreateWallet(username string) string {
	response := s.data.CreateWallet(username)
	return response
}
func (s *WalletService) OneUserWallet(username string) (int, bool) {
	balance, ok := s.data.GetWalletByName(username)
	if ok {
		return balance, ok
	}
	return balance, ok
}

func (s *WalletService) ShowAllWallets() map[string]int {
	response := s.data.GetWallets()
	return response
}

func (s *WalletService) UpdateWallet(username string, value int) string {
	response := s.data.WalletTransaction(username, value)
	return response
}

func NewService(data data.IData) *WalletService {
	return &WalletService{data: data}
}
