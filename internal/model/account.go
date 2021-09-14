package model

type Account struct {
	Name     string
	Broker   string
	TailFour string
	Stocks   []*Stock
}
