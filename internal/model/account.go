package model

type Account struct {
	ID       int64
	Name     string
	Broker   string
	TailFour string
	Stocks   []*Stock
}
