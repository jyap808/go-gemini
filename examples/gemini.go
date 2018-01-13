package main

import (
	"fmt"

	"github.com/jyap808/go-gemini"
)

const (
	API_KEY    = ""
	API_SECRET = ""
)

func main() {
	// Gemini client
	gemini := gemini.New(API_KEY, API_SECRET)

	// Get Ticker (btcusd)
	ticker, err := gemini.GetTicker("btcusd")
	fmt.Println(err, ticker)
	fmt.Println("Last:", ticker.Last)

	// Get Distribution (JBS)
	/* distribution, err := gemini.GetDistribution("SPHR")
	    fmt.Println(err)
		for _, balance := range distribution.Distribution {
			fmt.Println(balance.BalanceD)
		} */

	// Get market summaries
	/*
		marketSummaries, err := gemini.GetMarketSummaries()
		fmt.Println(err, marketSummaries)
	*/

	// Get orders book
	/*
		orderBook, err := gemini.GetOrderBook("BTC-DRK", "both", 100)
		fmt.Println(err, orderBook)
	*/

	// Market history
	/*
		marketHistory, err := gemini.GetMarketHistory("BTC-DRK", 100)
		for _, trade := range marketHistory {
			fmt.Println(err, trade.Timestamp.String(), trade.Quantity, trade.Price)
		}
	*/

	// Market

	// BuyLimit
	/*
		uuid, err := gemini.BuyLimit("BTC-DOGE", 1000, 0.00000102)
		fmt.Println(err, uuid)
	*/

	// BuyMarket
	/*
		uuid, err := gemini.BuyLimit("BTC-DOGE", 1000)
		fmt.Println(err, uuid)
	*/

	// Sell limit
	/*
		uuid, err := gemini.SellLimit("BTC-DOGE", 1000, 0.00000115)
		fmt.Println(err, uuid)
	*/

	// Cancel Order
	/*
		err := gemini.CancelOrder("e3b4b704-2aca-4b8c-8272-50fada7de474")
		fmt.Println(err)
	*/

	// Get open orders
	/*
		orders, err := gemini.GetOpenOrders("BTC-DOGE")
		fmt.Println(err, orders)
	*/

	// Account
	// Get balances
	/*
		balances, err := gemini.GetBalances()
		fmt.Println(err, balances)
	*/

	// Get balance
	/*
		balance, err := gemini.GetBalance("DOGE")
		fmt.Println(err, balance)
	*/

	// Get address
	/*
		address, err := gemini.GetDepositAddress("QBC")
		fmt.Println(err, address)
	*/

	// WithDraw
	/*
		whitdrawUuid, err := gemini.Withdraw("QYQeWgSnxwtTuW744z7Bs1xsgszWaFueQc", "QBC", 1.1)
		fmt.Println(err, whitdrawUuid)
	*/

	// Get order history
	/*
		orderHistory, err := gemini.GetOrderHistory("BTC-DOGE", 10)
		fmt.Println(err, orderHistory)
	*/

	// Get getwithdrawal history
	/*
		withdrawalHistory, err := gemini.GetWithdrawalHistory("all", 0)
		fmt.Println(err, withdrawalHistory)
	*/

	// Get deposit history
	/*
		deposits, err := gemini.GetDepositHistory("all", 0)
		fmt.Println(err, deposits)
	*/

}
