package main

import (
	"bootcamp/config"
	"bootcamp/data"
	"bootcamp/handler"
	"bootcamp/service"
	"fmt"
	"net/http"
)

func main() {
	config := config.Get()

	data := data.NewData(config.InitialBalanceAmount, config.MinimumBalanceAmount)
	service := service.NewService(data)
	h := handler.NewHandler(service)
	http.HandleFunc("/", h.WalletHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
