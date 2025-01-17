package handlers

import (
	"fmt"
	"net/http"
)

func GetInrBalances(req *http.Request) {
	fmt.Println("Get INR balance")
}

func GetStockBalances(req *http.Request) {
	fmt.Println("Get stock balance")
}

func OnRamp(req *http.Request) {
	fmt.Println("Onramp INR -> BTC")
}
