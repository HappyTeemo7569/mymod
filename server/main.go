package server

import (
	"github.com/HappyTeemo7569/mymod/base"
	"github.com/HappyTeemo7569/mymod/server/define"
	"github.com/HappyTeemo7569/mymod/server/myHttp"
	"github.com/HappyTeemo7569/mymod/server/myWebscoket"
)

func Run() {

	go myWebscoket.StartWebSocket()
	base.Logger.Infof("webSocket服务启动")

	go myHttp.RunHttp()
	base.Logger.Infof("htttp服务启动")

	<-define.CloseChan
	base.Logger.Info("收到关闭消息")
}
