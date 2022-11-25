package portfolio

import (
	"encoding/json"
	"fmt"

	"github.com/forsunforson/profolio/internal/model"
	"gorm.io/datatypes"
)

type PortfolioPO struct {
	ID          int64
	Accounts    datatypes.JSON
	Holders     datatypes.JSON
	Stocks      datatypes.JSON
	Cash        int
	MarketValue int
	Total       int
}

func (po PortfolioPO) TableName() string {
	return "portfolio_info"
}

func (po PortfolioPO) GetModel() model.Portfolio {
	var accounts []int64
	if err := json.Unmarshal([]byte(po.Accounts.String()), &accounts); err != nil {
		fmt.Printf("unmarshal accounts column failed, got %s, err %v", po.Accounts.String(), err)
	}
	var holders map[int64]*model.Holder
	if err := json.Unmarshal([]byte(po.Holders.String()), &holders); err != nil {
		fmt.Printf("unmarshal holders column failed, got %s, err %v", po.Holders.String(), err)
	}
	var stocks []string
	if err := json.Unmarshal([]byte(po.Stocks.String()), &stocks); err != nil {
		fmt.Printf("unmarshal stocks column failed, got %s, err %v", po.Stocks.String(), err)
	}
	return model.Portfolio{
		Total:       po.Total,
		Cash:        po.Cash,
		MarketValue: po.MarketValue,
		Accounts:    accounts,
		Holders:     holders,
		Stocks:      stocks,
		ID:          po.ID,
	}
}

func GetPO(p *model.Portfolio) PortfolioPO {
	var accounts datatypes.JSON
	if raw, err := json.Marshal(p.Accounts); err == nil {
		accounts.UnmarshalJSON(raw)
	}
	var holders datatypes.JSON
	if raw, err := json.Marshal(p.Holders); err == nil {
		holders.UnmarshalJSON(raw)
	}
	var stocks datatypes.JSON
	if raw, err := json.Marshal(p.Stocks); err == nil {
		stocks.UnmarshalJSON(raw)
	}
	return PortfolioPO{
		Total:       p.Total,
		MarketValue: p.MarketValue,
		Cash:        p.Cash,
		Accounts:    accounts,
		Holders:     holders,
		Stocks:      stocks,
		ID:          p.ID,
	}
}
