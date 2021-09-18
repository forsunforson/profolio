package cmd

import (
	"fmt"

	"github.com/forsunforson/profolio/internal/logic"
	"github.com/golang/glog"
)

func portfolioInterface() {
	for {
		fmt.Println("------ 组合管理 ------")
		fmt.Println("\t0. 查看所有组合")
		fmt.Println("\t1. 查看[i]组合信息")
		fmt.Println("\t2. 查看[i]组合股东信息")
		fmt.Println("\t3. 管理[i]组合股票信息")
		fmt.Println("\t4. 增加组合")
		fmt.Println("\t9. 返回上一级")
		var n string
		fmt.Scan(&n)
		switch n {
		case "0":
			showAllPortfolios(10, 0)
		case "1":
			getPortfolioInfo()
		case "2":
			getHolderInfo()
		case "3":
			managePortStock()
		case "4":
			addNewPort()
		case "9":
			return
		default:
			fmt.Println("请输入正确的数字")
		}
	}
}

func managePortStock() {
	fmt.Print("输入组合ID：")
	var pID int
	fmt.Scan(&pID)
	if !logic.IsInvalidPort(pID) {
		fmt.Printf("组合ID[%d]不存在\n", pID)
		return
	}

}

func getHolderInfo() {
	fmt.Print("输入组合ID：")
	var pID int
	fmt.Scan(&pID)
	port, err := logic.GetPortfolioInfo(pID)
	if err != nil {
		fmt.Println("查询不到这个ID，请重试")
		return
	}
	holders, err := logic.GetAllHolders(pID)
	if err != nil {
		fmt.Println("查询不到相关股东，请重试")
		return
	}
	if len(holders) == 0 {
		fmt.Println("该组合下无股东")
		return
	}
	for _, holder := range holders {
		value := int(holder.Percentage * float32(port.Total))
		fmt.Printf("组合[%d] 股东[%s] 持有份额[%.2f%%] 价值[%d]\n", pID, holder.Name, holder.Percentage*100, value)
	}
}

func getPortfolioInfo() {

}

func addNewPort() {
	var market, cash int
	fmt.Print("输入组合市值：")
	fmt.Scan(&market)
	fmt.Print("输入组合现金：")
	fmt.Scan(&cash)
	fmt.Print("确认创建？[Y/n]")
	var confirm string
	fmt.Scan(&confirm)
	switch confirm {
	case "y", "Y":
		break
	case "n", "N":
		return
	}
	idx, err := logic.AddNewPortfolio(market, cash)
	if err != nil {
		fmt.Println("生成新的组合失败，请重试")
		return
	}
	fmt.Printf("生成新的组合成功，你的组合ID为[%d]\n", idx)
}

func showAllPortfolios(pageSize, page int) {
	ports, err := logic.GetAllPortfolios(pageSize, page)
	if err != nil {
		fmt.Println("发生了一些错误，请重试")
		glog.Errorf("get all portfolios fail: %s", err)
		return
	}
	for _, port := range ports {
		fmt.Printf("组合编号[%d] 总金额[￥%d] 市值[￥%d] 现金[￥%d]\n", port.ID, port.Total, port.MarketValue, port.Cash)
	}
	if len(ports) < 10 {
		fmt.Println("---------- 已到达末尾 ----------")
		fmt.Println("按任意键退回上一级。。。。")
		var n int
		fmt.Scanln(&n)
		return
	}

	fmt.Println("翻下一页[n] 翻上一页[u] 任意键结束")
	var n string
	fmt.Scan(&n)
	switch n {
	case "n":
		showAllPortfolios(pageSize, page+1)
	case "u":
		if page > 1 {
			showAllPortfolios(pageSize, page-1)
		} else {
			fmt.Println("---------- 已到达首页 ----------")
			showAllPortfolios(pageSize, 0)
		}

	default:
		return
	}
}
