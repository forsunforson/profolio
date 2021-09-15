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
	return nil, errors.New("portfolio not found")
}

func GetAllPortfolios(pageSize, page int) ([]*model.Portfolio, error) {
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

func AddNewPortfolio(total, market, cash int) (int, error) {
	db := pool.Database
	sql := "insert into portfolio_info (total_value, market_value, cash) values (?,?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		glog.Errorf("prepare sql[%s] fail: %s\n", sql, err)
		return -1, err
	}
	ret, err := stmt.Exec(total, market, cash)
	if err != nil {
		glog.Errorf("exec sql fail: %s\n", err)
		return -1, err
	}
	idx, err := ret.LastInsertId()
	if err != nil {
		glog.Errorf("get idx fail: %s", err)
		return -1, nil
	}
	return int(idx), nil
}
