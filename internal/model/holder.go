package model

type Holder struct {
	Name       string
	Total      int
	Percentage float32
	portfolio  *Portfolio
}

func (h *Holder) Buy(money int) {
	h.portfolio.Lock.Lock()
	defer h.portfolio.Lock.Unlock()
	h.portfolio.Total = h.portfolio.Total + money
	h.Total += money

	for _, holder := range h.portfolio.Holders {
		holder.Percentage = float32(holder.Total) / float32(h.portfolio.Total)
	}
}

func (h *Holder) Sell(money int) {
	h.portfolio.Lock.Lock()
	defer h.portfolio.Lock.Unlock()
	h.portfolio.Total = h.portfolio.Total - money
	h.Total -= money

	for _, holder := range h.portfolio.Holders {
		holder.Percentage = float32(holder.Total) / float32(h.portfolio.Total)
	}
}
