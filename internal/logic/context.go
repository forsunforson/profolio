package logic

import (
	"github.com/forsunforson/profolio/internal/model"
)

type RuntimeContext struct {
	Portfolios []*model.Portfolio
	stocks     []model.Stock
}

var (
	runtimeContext *RuntimeContext
)

func InitContext() {
	ctx := RuntimeContext{}
	stocks := GetAllStocks()
	ctx.stocks = stocks
	runtimeContext = &ctx
	go Ticker()
}

func GetRunTimeContext() *RuntimeContext {
	return runtimeContext
}
