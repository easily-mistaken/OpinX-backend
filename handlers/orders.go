package handlers

import (
	"fmt"
	"net/http"
)

func BuyOrder(req *http.Request) {
	fmt.Println("Buy order")
}

func SellOrder(req *http.Request) {
	fmt.Println("Sell order")
}

func ViewOrder(req *http.Request) {
	fmt.Println("View order")
}
