package cmd

import (
	"fmt"

	"github.com/forsunforson/profolio/internal/logic"
)

func CommandReceiver() {
	for {
		fmt.Println("输入数字，选择操作模块:")
		fmt.Println("\t0. 管理界面")
		fmt.Println("\t1. 展示界面")
		fmt.Println("\t9. 退出程序")
		var n string
		_, _ = fmt.Scan(&n)
		switch n {
		case "0":
			showManageInterface()
		case "1":
			showDisplayInterface()
		case "9":
			return
		default:
			fmt.Printf("请输入正确的数字\n")
		}
	}
}

func showManageInterface() {
	for {
		fmt.Println("------ 管理模块 ------")
		fmt.Println("\t0. 组合管理")
		fmt.Println("\t1. 持有人管理")
		fmt.Println("\t2. 账户管理")
		fmt.Println("\t3. 股票管理")
		fmt.Println("\t9. 返回上一级")
		var n string
		_, _ = fmt.Scan(&n)
		switch n {
		case "0":
			portfolioInterface()
		case "9":
			return
		default:
			fmt.Println("请输入正确的数字")
		}
	}
}

func showDisplayInterface() {

}

func addHolder() {
	fmt.Printf("Please Input A Name: \n")
	var name string
	_, _ = fmt.Scan(&name)
	err := logic.NewHolder(name, 0)
	if err != nil {
		fmt.Printf("Add Holder fail, maybe there is a same name or please try again.\n")
		return
	}
	fmt.Printf("Add Holder %s successfully.\n", name)
}

func showMarketValue() {
	ctx := logic.GetRunTimeContext()
	for _, holder := range ctx.Portfolios[0].Holders {
		fmt.Printf("%s has %.2f%% , value %d yuan\n", holder.Name, holder.Percentage*100, holder.Total)
	}
}
