package service

import (
	"fmt"

	"github.com/forsunforson/profolio/internal/repository/portfolio"
	"github.com/forsunforson/profolio/internal/repository/user"
)

type TradeService interface {
	BuyIn(portfolioID, userID int64, cash int)
	SellOut(portfolioID, userID int64, amount int)
}

type tradeServiceImpl struct {
	userRepo      user.Repository
	portfolioRepo portfolio.Repository
}

func NewTradeService() TradeService {
	service := tradeServiceImpl{
		userRepo:      user.NewRepository(),
		portfolioRepo: portfolio.NewRepository(),
	}
	return &service
}

func (i *tradeServiceImpl) BuyIn(portfolioID, userID int64, cash int) {
	portfolio, err := i.portfolioRepo.GetPortfolioByID(portfolioID)
	if err != nil {
		fmt.Printf("get profolio failed: %v\n", err)
		return
	}
	user, err := i.userRepo.GetUserByID(userID)
	if err != nil {
		fmt.Printf("get user failed: %v\n", err)
		return
	}
	holder, err := portfolio.BuyIn(userID, cash)
	if err != nil {
		fmt.Printf("buy in ￥%d for %s fail", cash, user.Name)
		return
	}
	fmt.Printf("user %s have %.2f%% in portforlio %d\n", user.Name, holder.Percentage*100, portfolioID)
}

func (i *tradeServiceImpl) SellOut(portfolioID, userID int64, amount int) {

}
