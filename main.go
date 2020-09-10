package main

import (
	"github.com/HappyTeemo7569/mymod/base"
	"github.com/HappyTeemo7569/mymod/server"
)

func main() {
	base.Logger.Info("模块加载完成")
	server.Run()
}
