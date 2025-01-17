package config

import "github.com/easily-mistaken/OpinX-backend/types"

// InitialINRBalances represents the initial INR balances for users
var InitialINRBalances = types.INRBalances{
	"user1": {
		Balance: 10,
		Locked:  0,
	},
	"user2": {
		Balance: 20,
		Locked:  10,
	},
}

// InitialOrderBook represents the initial order book state
var InitialOrderBook = types.OrderBook{
	"BTC_USDT_10_Oct_2024_9_30": {
		Yes: map[types.PriceRange]types.OrderData{
			types.Price9_5: {
				Total: 12,
				Orders: []types.Order{
					{UserID: "user1", Quantity: 2},
					{UserID: "user2", Quantity: 10},
				},
			},
			types.Price8_5: {
				Total: 12,
				Orders: []types.Order{
					{UserID: "user1", Quantity: 3},
					{UserID: "user2", Quantity: 3},
					{UserID: "user3", Quantity: 6},
				},
			},
		},
		No: make(map[types.PriceRange]types.OrderData),
	},
}

// InitialStockBalances represents the initial stock balances for users
var InitialStockBalances = types.StockBalances{
	"user1": {
		"BTC_USDT_10_Oct_2024_9_30": {
			Yes: &types.StockPosition{
				Quantity: 1,
				Locked:   0,
			},
		},
	},
	"user2": {
		"BTC_USDT_10_Oct_2024_9_30": {
			No: &types.StockPosition{
				Quantity: 3,
				Locked:   4,
			},
		},
	},
}
  