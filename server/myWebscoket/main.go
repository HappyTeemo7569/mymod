package myWebscoket

import (
	"encoding/json"
	"fmt"
	"github.com/HappyTeemo7569/mymod/base"
	"github.com/HappyTeemo7569/mymod/server/define"
	"golang.org/x/net/websocket"
	"net/http"
	"time"
)

func init() {
	//创建消息队列
	NormalChan = make(QueChan, 500000)
	go queueNormal()
}

func StartWebSocket() {
	pattern := fmt.Sprintf("/%s", base.ConfigServer.WsName)

	http.Handle(pattern, websocket.Handler(WS_Thread))
	base.Logger.Infof("启动webSocket:", pattern)

	go func() {
		addr := fmt.Sprintf(":%d", base.ConfigServer.WsPort)
		base.Logger.Infof("启动webSocket:监听:", addr)
		if err := http.ListenAndServe(addr, nil); err != nil {
			base.Logger.Errorf("websocket启动失败！", err)
			define.CloseChan <- 0
		} else {
			base.Logger.Infof("开始监听", addr)
		}
	}()
}

func WS_Thread(ws *websocket.Conn) {
	NormalChan <- WsMsgItem{CONN_TYPE_CONNECT, ws, nil}

	var err error
	for {
		var msg string

		if err = websocket.Message.Receive(ws, &msg); err != nil {
			base.Logger.Errorf("接收消息错误：", err)
			NormalChan <- WsMsgItem{CONN_TYPE_CLOSE, ws, nil}
			return
		}

		skt := GetSocketItem(ws)
		if skt == nil {
			base.Logger.Errorf("未保存的连接")
			NormalChan <- WsMsgItem{CONN_TYPE_CLOSE, ws, nil}
			return
		}

		var WSdata TagClient_Universal
		if err = json.Unmarshal([]byte(msg), &WSdata); err != nil {
			str, _ := json.Marshal(WSdata)
			base.Logger.Errorf("解析消息错误：", err, string(str))
			NormalChan <- WsMsgItem{CONN_TYPE_CLOSE, ws, nil}
			return
		}

		if WSdata.InterfaceNumber == HEARTBEAT_MSG_ID { //心跳包
			if skt.UserData == nil {
				NormalChan <- WsMsgItem{CONN_TYPE_CLOSE, ws, nil}
				return
			}
			skt.MsgTime = int(time.Now().Unix()) //更新时间戳
			sendMsg := &tagServer_HeartBeat{
				HEARTBEAT_MSG_ID,
			}
			WS_SendMessage(sendMsg, ws)
		} else {
			NormalChan <- WsMsgItem{CONN_TYPE_READ, ws, &WSdata}
		}

	}
}

func WS_OnConnect(ws *websocket.Conn) {
	skt := GetSocketItem(ws)
	if skt != nil {
		base.Logger.Warningf("重复连接...", ws)
	} else {
		AddSocketItem(ws)
	}
}

func WS_Close(ws *websocket.Conn) {
	ws.Close()
	skt := GetSocketItem(ws)
	if skt == nil {
		base.Logger.Warningf("断开连接，但是没找到连接...", ws)
		return
	}
	define.ExitChan <- define.ExitGameMsgItem{skt.UserData}
	DelSocketItem(ws)
}
