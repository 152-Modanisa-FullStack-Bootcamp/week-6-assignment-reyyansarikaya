package data

type IData interface {
	GetWallets() map[string]int
	GetWalletByName(string) (int, bool)
	WalletTransaction(string, int) string
	CreateWallet(string) string
}

var Storage = make(map[string]int)

type Data struct {
	Initialbalance int
	Minimumbalance int
}

func (Data) GetWallets() map[string]int {
	return Storage
}

func (d Data) CreateWallet(username string) string {
	_, ok := Storage[username]
	if !ok {
		Storage[username] = d.Initialbalance
		return "Wallet created"
	}
	return "Username already created"
}

func (Data) GetWalletByName(username string) (int, bool) {

	balance, ok := Storage[username]
	return balance, ok
}
func (d Data) WalletTransaction(username string, value int) string {
	balance, ok := Storage[username]
	if ok {
		newBalance := balance + value
		if newBalance < d.Minimumbalance {
			return "The transaction could not be processed because there is no balance in the wallet."
		} else {
			Storage[username] = newBalance
			return "Transaction successful."
		}
	}
	return "Username not found."

}
func NewData(initialbalance int, minimumbalance int) IData {
	return &Data{Initialbalance: initialbalance, Minimumbalance: minimumbalance}
}
