package myWebscoket

import (
	"github.com/HappyTeemo7569/mymod/server/define"
)

//一般队列
func queueNormal() {
	for {
		data := <-NormalChan

		if data.Conntype != CONN_TYPE_CHECK {
			//base.Logger.Debugf("收到一般消息:conntype:", data.Conntype, ":数据:", string(mes), ":消息队列长度:", len(Normalchan))
		}

		switch data.Conntype {
		case CONN_TYPE_CONNECT:
			WS_OnConnect(data.Ws)
		case CONN_TYPE_CLOSE:
			WS_Close(data.Ws)
		case CONN_TYPE_READ: //业务逻辑
			define.JoinChan <- define.JoinMsgItem{data.Ws, data.WSdata}
		case CONN_TYPE_CHECK:
			CheckSocketList()
		}

		if data.Conntype != CONN_TYPE_CHECK {
			//base.Logger.Debugf("一般消息处理完成:conntype:", data.Conntype, ":数据:", string(mes), ":消息队列长度:", len(Normalchan))
		}
	}
}
