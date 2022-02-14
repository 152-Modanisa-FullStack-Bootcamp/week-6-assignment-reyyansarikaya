package service

import (
	"bootcamp/config"
	"bootcamp/data"
	"reflect"
	"testing"
)

func TestWalletService_CreateWallet(t *testing.T) {
	type fields struct {
		data data.IData
	}
	tests := []struct {
		name     string
		fields   fields
		username string
		want     string
	}{
		{"case1", fields{data: data.NewData(config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount)}, "reyyan", "Wallet created"},
		{"case1", fields{data: data.NewData(config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount)}, "reyyan", "Username already created"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WalletService{
				data: tt.fields.data,
			}
			if got := s.CreateWallet(tt.username); got != tt.want {
				t.Errorf("CreateWallet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletService_ShowAllWallets(t *testing.T) {
	response := make(map[string]int)
	response["reyyan"] = 0
	type fields struct {
		data data.IData
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]int
	}{
		{"case1", fields{data: data.NewData(config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount)}, response},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WalletService{
				data: tt.fields.data,
			}
			if got := s.ShowAllWallets(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShowAllWallets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalletService_OneUserWallet(t *testing.T) {
	type fields struct {
		data data.IData
	}
	tests := []struct {
		name     string
		fields   fields
		username string
		want     int
		want1    bool
	}{
		{"case1", fields{data: data.NewData(config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount)}, "reyyan", 0, true},
		{"case1", fields{data: data.NewData(config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount)}, "fatma", 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WalletService{
				data: tt.fields.data,
			}
			got, got1 := s.OneUserWallet(tt.username)
			if got != tt.want {
				t.Errorf("OneUserWallet() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("OneUserWallet() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWalletService_UpdateWallet(t *testing.T) {
	type fields struct {
		data data.IData
	}
	tests := []struct {
		name     string
		fields   fields
		username string
		value    int
		want     string
	}{
		{"case1", fields{data: data.NewData(config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount)}, "reyyan", 100, "Transaction successful."},
		{"case2", fields{data: data.NewData(config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount)}, "reyyan", -1000, "The transaction could not be processed because there is no balance in the wallet."},
		{"case3", fields{data: data.NewData(config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount)}, "fatma", 100, "Username not found."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WalletService{
				data: tt.fields.data,
			}
			if got := s.UpdateWallet(tt.username, tt.value); got != tt.want {
				t.Errorf("UpdateWallet() = %v, want %v", got, tt.want)
			}
		})
	}
}
