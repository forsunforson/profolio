package cmd

import (
	"fmt"

	"github.com/forsunforson/profolio/internal/logic"
	"github.com/forsunforson/profolio/internal/model"
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
		case "3":
			addNewStock()
		case "9":
			return
		default:
			fmt.Println("请输入正确的数字")
		}
	}
}

func addNewStock() {
	fmt.Println("正在添加股票......")
	fmt.Println("请选择股票市场:")
	fmt.Println("\t1.沪深\n\t2.香港\n\t3.美国\n\t其他任意键退出")
	var n string
	fmt.Scan(&n)
	market := ""
	prefix := ""
	switch n {
	case "1":
		market = model.JuheHushen
		fmt.Println("请选择:\n\t1.沪市\n\t2.深市\n\t其他任意键退出")
		fmt.Scan(&n)
		if n == "1" {
			prefix = "sh"
		} else if n == "2" {
			prefix = "sz"
		} else {
			return
		}
	case "2":
		market = model.JuheHongkong
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
	stock, err := model.NewStock(prefix+code, market)
	if err != nil {
		fmt.Printf("找不到该股票或者发生错误：%s\n", err)
		return
	}
	// TODO 使用异步通知完成
	go logic.AddNewStock(stock)
	fmt.Println("成功增加股票：" + stock.GetName())
}
