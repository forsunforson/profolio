package model

import "sync"

var ()

type Portfolio struct {
	Accounts    []*Account
	Holders     []*Holder
	Cash        int
	MarketValue int
	Total       int
	ID          int

	Lock sync.RWMutex
}
