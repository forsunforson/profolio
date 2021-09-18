package logic

import (
	"errors"
	"fmt"

	"github.com/forsunforson/profolio/internal/model"
	"github.com/forsunforson/profolio/internal/pool"
	"github.com/golang/glog"
)

func PortfolioManager() {

}

// IsInvalidPort 检查组合ID是否有效
func IsInvalidPort(id int) bool {
	for _, p := range runtimeContext.Portfolios {
		if p.ID == id {
			return true
		}
	}
	return false
}

// GetPortfolioInfo 根据组合ID获得组合的相关信息，包括净值和持仓
func GetPortfolioInfo(id int) (*model.Portfolio, error) {
	db := pool.Database
	sql := "select * from portfolio_info where id = ? limit 10"
	rows, err := db.Query(sql, id)
	if err != nil {
		glog.Errorf("query sql[%s] fail: %s\n", sql, err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, marketValue, cash, totalValue int
		err := rows.Scan(&id, &marketValue, &cash, &totalValue)
		if err != nil {
			glog.Errorf("read row fail: %s", err)
			continue
		}
		p := &model.Portfolio{
			ID:          id,
			MarketValue: marketValue,
			Cash:        cash,
			Total:       totalValue,
		}
		return p, nil
	}
	// TODO 从portfolio_stock中读取组合的持仓
	return nil, errors.New("portfolio not found")
}

// GetAllPortfolios 按页读取所有的组合信息
func GetAllPortfolios(pageSize, page int) ([]*model.Portfolio, error) {
	// TODO 后期先从runtime读，再从数据库读
	db := pool.Database
	portfolios := make([]*model.Portfolio, 0)
	sql := fmt.Sprintf("select * from portfolio_info order by id limit 10 offset %d", page)
	rows, err := db.Query(sql)
	if err != nil {
		glog.Errorf("query sql[%s] fail: %s\n", sql, err)
		return nil, err
	}
	defer rows.Close()
	counter := 0
	for rows.Next() {
		counter++
		var id, marketValue, cash, totalValue int
		rows.Scan(&id, &marketValue, &cash, &totalValue)
		p := &model.Portfolio{
			ID:          id,
			MarketValue: marketValue,
			Cash:        cash,
			Total:       totalValue,
		}
		portfolios = append(portfolios, p)
	}

	return portfolios, nil
}

func AddNewPortfolio(marketValue, cash int) (int, error) {
	db := pool.Database
	sql := "insert into portfolio_info (total_value, market_value, cash) values (?,?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		glog.Errorf("prepare sql[%s] fail: %s\n", sql, err)
		return -1, err
	}
	ret, err := stmt.Exec(marketValue, cash)
	if err != nil {
		glog.Errorf("exec sql fail: %s\n", err)
		return -1, err
	}

	idx, err := ret.LastInsertId()
	if err != nil {
		glog.Errorf("get idx fail: %s", err)
		return -1, nil
	}
	runtimeContext.Portfolios = append(runtimeContext.Portfolios, &model.Portfolio{
		MarketValue: marketValue,
		Cash:        cash,
		Total:       marketValue + cash,
		ID:          int(idx),
	})
	return int(idx), nil
}
