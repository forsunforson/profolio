package logic

import (
	"fmt"
	"time"

	"github.com/forsunforson/profolio/internal/model"
	"github.com/forsunforson/profolio/internal/pool"
	"github.com/golang/glog"
)

// Ticker 定时任务，只要服务在运行就会更新股票价格
func Ticker() {
	// 开始运行程序时，先更新尝试着更新一下
	tick := time.NewTicker(1 * time.Second)
	for {
		<-tick.C
		if time.Now().Local().Hour() > 16 {
			// 收盘了来更新一下
			if runtimeContext.lastUpdate.Day() != time.Now().Day() {
				// 更新所有关注的股票价格
				UpdateAllStocks()
				runtimeContext.lastUpdate = time.Now()
				// TODO更新所有组合的净值
				// TODO更新所有股东的净值
			}

			tick.Reset(1 * time.Hour)
		}
	}
}

func GetAllStocks() []model.Stock {
	ctx := GetRunTimeContext()
	stocks := make([]model.Stock, 0)
	for _, v := range ctx.Stocks {
		stocks = append(stocks, v)
	}
	return stocks
}

// readAllStocks 读取组合中所有的股票，加入需要监听价格的队列
func readAllStocks() []model.Stock {
	stocks := make(map[string]*model.JuheStock)
	db := pool.Database
	sql := "select * from stock_info order by id"
	row, err := db.Query(sql)
	if err != nil {
		glog.Errorf("query sql[%s] fail: %s", sql, err)
		return nil
	}
	for row.Next() {
		var id int
		var mtime time.Time
		var stockCode, stockName, market string
		err := row.Scan(&id, &mtime, &stockCode, &stockName, &market)
		if err != nil {
			glog.Errorf("read row fail: %s", err)
			continue
		}
		stock := &model.JuheStock{
			ID:        id,
			Mtime:     mtime,
			StockCode: stockCode,
			StockName: stockName,
			Market:    market,
		}
		stocks[stock.StockCode] = stock
	}
	row.Close()
	stockList := make([]model.Stock, 0)
	for code, stock := range stocks {
		sql = "select * from stock_price where stock_id = ? order by cdate desc limit 1"
		row, err = db.Query(sql, code)
		if err != nil {
			continue
		}
		for row.Next() {
			var stockID, cdate, price string
			err := row.Scan(&stockID, &cdate, &price)
			if err != nil {
				continue
			}
			stock.Price = price
			stock.CDate = cdate
			stockList = append(stockList, stock)
		}

		row.Close()
	}

	return stockList
}

func UpdateAllStocks() {
	stocks := GetRunTimeContext().Stocks
	for code, stock := range stocks {
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
		stocks[code] = stock
	}
}

// AddNewStock TODO 实现写入数据库逻辑
func AddNewStock(stock model.Stock) (int, error) {
	db := pool.Database
	sql := "insert into stock_info (stock_code, stock_name, market) values (?,?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		glog.Errorf("prepare sql[%s] fail: %s\n", sql, err)
		return -1, err
	}
	ret, err := stmt.Exec(stock.GetCode(), stock.GetName(), stock.GetMarket())
	if err != nil {
		glog.Errorf("exec sql fail: %s\n", err)
		return -1, err
	}
	// 加入运行环境变量
	runtimeContext.Stocks[stock.GetCode()] = stock

	idx, err := ret.LastInsertId()
	if err != nil {
		glog.Errorf("get idx fail: %s", err)
		return -1, nil
	}
	return int(idx), nil
}
