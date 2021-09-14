package logic

import (
	"fmt"
	"time"

	"github.com/forsunforson/profolio/internal/model"
	"github.com/forsunforson/profolio/internal/pool"
	"github.com/golang/glog"
)

func init() {
	go ticker()
}

func ticker() {
	// 定时任务，只要服务在运行就会更新股票价格
	tick := time.NewTicker(10 * time.Hour)
	for {
		<-tick.C
		if time.Now().Local().Hour() > 16 {
			UpdateStock()
		}
	}
}

func GetAllStocks(stocksCode []string) []model.Stock {
	stocks := make([]model.Stock, 0)
	for _, code := range stocksCode {
		stock, err := model.NewStock(code, model.JuheHongkong)
		if err != nil {
			continue
		}
		stocks = append(stocks, stock)
	}
	return stocks
}

func UpdateStock() {
	stocks := GetRunTimeContext().stocks
	for _, stock := range stocks {
		now := time.Now()
		db := pool.Database

		date := fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())
		sql := "select count(*) from stock_price where stock_id = ? and cdate = ?"
		row, err := db.Query(sql, stock.GetCode(), date)
		if err != nil {
			glog.Errorf("query sql[%s] fail: %s", sql, err)
			continue
		}
		if row.Next() {
			var count int
			row.Scan(&count)
			if count > 0 {
				glog.Infof("stock_id[%s] cdate[%s] exist", stock.GetCode(), date)
				row.Close()
				continue
			}
		}
		row.Close()
		stock, err := model.NewStock(stock.GetCode()[2:], model.JuheHongkong)
		if err != nil {
			continue
		}
		sql = "insert into stock_price (stock_id, price) values (?,?)"
		stmt, err := db.Prepare(sql)
		if err != nil {
			glog.Errorf("prepare sql[%s] fail: %s\n", sql, err)
			continue
		}
		_, err = stmt.Exec(stock.GetCode(), stock.GetLatestPrice())
		if err != nil {
			glog.Errorf("exec sql[%s] fail: %s\n", sql, err)
			continue
		}
	}

}
