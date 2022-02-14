package data

import (
	"bootcamp/config"
	"reflect"
	"testing"
)

func TestData_CreateWallet(t *testing.T) {
	response := NewData(config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount)
	response.CreateWallet("reyyan")
	type fields struct {
		Initialbalance int
		Minimumbalance int
	}
	tests := []struct {
		name     string
		username string
		fields   fields
		want     string
	}{
		{"case1", "fatma", fields{config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount}, "Wallet created"},
		{"case2", "reyyan", fields{config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount}, "Username already created"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				Initialbalance: tt.fields.Initialbalance,
				Minimumbalance: tt.fields.Minimumbalance,
			}
			if got := d.CreateWallet(tt.username); got != tt.want {
				t.Errorf("CreateWallet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_GetWalletByName(t *testing.T) {
	type fields struct {
		Initialbalance int
		Minimumbalance int
	}
	tests := []struct {
		name     string
		fields   fields
		username string
		want     int
		want1    bool
	}{
		{"case1", fields{config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount}, "ayse", 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := Data{
				Initialbalance: tt.fields.Initialbalance,
				Minimumbalance: tt.fields.Minimumbalance,
			}
			got, got1 := da.GetWalletByName(tt.username)
			if got != tt.want {
				t.Errorf("GetWalletByName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetWalletByName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestData_GetWallets(t *testing.T) {
	a := make(map[string]int)
	a["fatma"] = 0
	a["reyyan"] = 0
	type fields struct {
		Initialbalance int
		Minimumbalance int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]int
	}{
		{"case1", fields{config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount}, a},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := Data{
				Initialbalance: tt.fields.Initialbalance,
				Minimumbalance: tt.fields.Minimumbalance,
			}
			if got := da.GetWallets(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWallets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_WalletTransaction(t *testing.T) {
	type fields struct {
		Initialbalance int
		Minimumbalance int
	}
	tests := []struct {
		name     string
		fields   fields
		username string
		value    int
		want     string
	}{ // fields{data: data.NewData(config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount)}
		{"case1", fields{config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount}, "sarikaya", 100, "Username not found."},
		{"case1", fields{config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount}, "reyyan", 100, "Transaction successful."},
		{"case1", fields{config.C.InitialBalanceAmount, config.C.MinimumBalanceAmount}, "reyyan", -100000, "The transaction could not be processed because there is no balance in the wallet."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				Initialbalance: tt.fields.Initialbalance,
				Minimumbalance: tt.fields.Minimumbalance,
			}
			if got := d.WalletTransaction(tt.username, tt.value); got != tt.want {
				t.Errorf("WalletTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
