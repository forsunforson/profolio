package cmd

import (
	"fmt"
)

func stockInterface() {
	for {
		fmt.Println("------ 股票管理 ------")
		fmt.Println("\t0. 查看所有股票")
		fmt.Println("\t1. 查看[i]股票信息")
		fmt.Println("\t2. 查看[i]最近五日收盘价")
		fmt.Println("\t3. 增加股票")
		fmt.Println("\t9. 返回上一级")
		var n string
		fmt.Scan(&n)
		switch n {
		case "0":
			showAllStocks()
		case "3":
			addNewStock()
		case "9":
			return
		default:
			fmt.Println("请输入正确的数字")
		}
	}
}

func showAllStocks() {
	// stocks := logic.GetAllStocks()
	// // TODO 这里编号会变，使用编号展示股票信息是不靠谱的
	// for idx := 0; idx < len(stocks); idx++ {
	// 	fmt.Printf("编号[%d] 股票[%s] 代码[%s]\n", idx, stocks[idx].GetName(), stocks[idx].GetCode())
	// }

}

func addNewStock() {
	fmt.Println("正在添加股票......")
	fmt.Println("请选择股票市场:")
	fmt.Println("\t1.沪深\n\t2.香港\n\t3.美国\n\t其他任意键退出")
	var n string
	fmt.Scan(&n)
	//market := ""
	prefix := ""
	switch n {
	case "1":
	case "2":

	default:
		return
	}
	var code string
	fmt.Println("请输入股票代码：\n例如688007或者00700")
	fmt.Scan(&code)
	if prefix == "sh" || prefix == "sz" {
		if len(code) != 6 {
			fmt.Println("沪深股票代码应为6位")
			return
		}
	} else {
		if len(code) != 5 {
			fmt.Println("港股代码应为5位")
			return
		}
	}
}
