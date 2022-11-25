package model

import (
	"fmt"
	"sync"
)

type Portfolio struct {
	Accounts    []int64
	Holders     map[int64]*Holder
	Stocks      []string
	Cash        int
	MarketValue int
	Total       int
	ID          int64

	lock sync.RWMutex
}

func (p *Portfolio) BuyIn(userID int64, money int) (*Holder, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	holder, ok := p.Holders[userID]
	if !ok {
		holder = p.NewHolder(userID)
	}
	holder.Buy(money)
	p.Total += money
	for _, h := range p.Holders {
		h.Percentage = float32(holder.Total) / float32(p.Total)
	}
	return holder, nil
}

func (p *Portfolio) SellOut(userID int64, money int) (*Holder, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	holder, ok := p.Holders[userID]
	if !ok {
		return nil, fmt.Errorf("user %d not exist", userID)
	}
	if money > holder.Total {
		money = holder.Total
	}
	holder.Sell(money)
	p.Total -= money
	for _, h := range p.Holders {
		h.Percentage = float32(holder.Total) / float32(p.Total)
	}
	return holder, nil
}

func (p *Portfolio) NewHolder(userID int64) *Holder {
	holder := Holder{
		UserID: userID,
	}
	p.Holders[userID] = &holder
	return &holder
}

func (p *Portfolio) CleanHolders() []int64 {
	cleanedHolders := []int64{}
	for id, holder := range p.Holders {
		if holder.Total == 0 {
			cleanedHolders = append(cleanedHolders, id)
		}
	}
	return cleanedHolders
}
