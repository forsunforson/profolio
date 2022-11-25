package model

type User struct {
	ID         int64
	Name       string
	Portfolios []int64
	Total      int
}
