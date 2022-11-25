package main

import (
	"flag"
	"fmt"

	"github.com/forsunforson/profolio/cmd"
	"github.com/forsunforson/profolio/config"
	"github.com/golang/glog"
)

func main() {
	fmt.Println("hello world")
	config.InitGlobalConfig()

	op := flag.String("m", "0", "mode to start server")
	flag.Parse()
	glog.Infof("aaa")
	defer glog.Flush()
	switch *op {
	case "0":
		cmd.CommandReceiver()
	default:
		fmt.Println("not support")
	}
	fmt.Println("goodbye world")
}
