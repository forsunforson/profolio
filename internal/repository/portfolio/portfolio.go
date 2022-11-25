package portfolio

import (
	"fmt"

	"github.com/forsunforson/profolio/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository interface {
	GetPortfolioByID(id int64) (*model.Portfolio, error)
}

type portfolioImpl struct {
	db *gorm.DB
}

var _ Repository = &portfolioImpl{}

func NewRepository() Repository {
	repo := portfolioImpl{}
	db, err := gorm.Open(sqlite.Open("aaa.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("connect db fail: %s\n", err)
	}
	repo.db = db
	return &repo
}

func (p *portfolioImpl) GetPortfolioByID(id int64) (*model.Portfolio, error) {
	var po PortfolioPO
	result := p.db.First(&po, id)
	if result.Error != nil {
		return nil, result.Error
	}
	model := po.GetModel()
	return &model, nil
}

func (p *portfolioImpl) Save(porfolio *model.Portfolio) error {
	po := GetPO(porfolio)
	err := p.db.Model(&PortfolioPO{}).Where("id = ?", po.ID).Updates(map[string]interface{}{
		"cash":         po.Cash,
		"market_value": po.MarketValue,
		"total":        po.Total,
		"accounts":     po.Accounts,
		"holders":      po.Holders,
		"stocks":       po.Stocks,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
