package logic

import (
	"errors"
	"strconv"
	"time"

	"github.com/forsunforson/profolio/internal/model"
	"github.com/forsunforson/profolio/internal/pool"
	"github.com/golang/glog"
)

// GetHolder 从db里读用户
func GetHolder(name string, portfolioID int) (*model.Holder, error) {
	db := pool.Database
	sql := "select * from holder_info where h_name = ? and h_portfolio = ? limit 100"
	rows, err := db.Query(sql, name, portfolioID)
	if err != nil {
		glog.Errorf("query sql[%s] fail: %s\n", sql, err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, total int
		var mtime, ctime time.Time
		var name, percentage string
		err := rows.Scan(&id, &mtime, &ctime, &name, &percentage, &total)
		if err != nil {
			glog.Errorf("read row fail: %s", err)
			continue
		}
		per, _ := strconv.ParseFloat(percentage, 32)
		h := model.Holder{
			Name:       name,
			Total:      total,
			Percentage: float32(per),
		}
		return &h, nil
	}

	return nil, errors.New(name + " not found")
}

// GetAllHolders 读取组合下所有的股东
func GetAllHolders(pID int) ([]model.Holder, error) {
	db := pool.Database
	sql := "select * from holder_info where h_portfolio = ? limit 100"
	rows, err := db.Query(sql, pID)
	if err != nil {
		glog.Errorf("query sql[%s] fail: %s\n", sql, err)
		return nil, err
	}
	defer rows.Close()
	holders := make([]model.Holder, 0)
	for rows.Next() {
		var id, total, pID int
		var mtime, ctime time.Time
		var name, percentage string
		rows.Scan(&id, &mtime, &ctime, &name, &percentage, &total, &pID)
		per, _ := strconv.ParseFloat(percentage, 32)
		h := model.Holder{
			Name:       name,
			Total:      total,
			Percentage: float32(per),
		}
		holders = append(holders, h)
	}
	return holders, nil
}

func NewHolder(name string, portfolioID int) error {
	// 暂时只有一个账户，就是0
	db := pool.Database
	sql := "insert into holder_info(h_name, h_percentage, h_total, h_portfolio) values (?,0,0,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		glog.Errorf("prepare sql[%s] fail: %s\n", sql, err)
		return err
	}
	_, err = stmt.Exec(name, portfolioID)
	if err != nil {
		glog.Errorf("exec sql[%s] fail: %s\n", sql, err)
		return err
	}
	return nil
}

func UpdateHoldersValue(pID int) error {
	port, err := GetPortfolioInfo(pID)
	if err != nil {
		return err
	}
	holders, err := GetAllHolders(pID)
	if err != nil {
		return err
	}
	db := pool.Database
	for _, holder := range holders {
		sql := "update holder_info set h_total = ? where h_portfolio = ? and h_name = ?"
		stmt, err := db.Prepare(sql)
		if err != nil {
			glog.Errorf("prepare sql[%s] fail: %s", sql, err)
			continue
		}
		newTotal := int(holder.Percentage * float32(port.Total))
		_, err = stmt.Exec(newTotal, pID, holder.Name)
		if err != nil {
			glog.Errorf("exec sql[%s] fail: %s", sql, err)
			continue
		}
	}
	return nil
}
