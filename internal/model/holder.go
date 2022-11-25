package model

type Holder struct {
	UserID     int64
	Total      int
	Percentage float32
}

func (h *Holder) Buy(money int) {
	h.Total += money
}

func (h *Holder) Sell(money int) {
	h.Total -= money
}
