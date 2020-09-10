package myWebscoket

import (
	"encoding/json"
	"github.com/HappyTeemo7569/mymod/base"
	"golang.org/x/net/websocket"
)

//单发消息
func WS_SendMessage(sendMsg interface{}, ws *websocket.Conn) bool {
	msg, err := json.Marshal(sendMsg)
	if err != nil {
		base.Logger.Errorf("格式化消息错误：", err)
		WS_SendError(-1, "异常，请重新登录", ws)
		return false
	}

	if err = websocket.Message.Send(ws, string(msg)); err != nil {
		base.Logger.Errorf("发送消息错误：", err)
		return false
	}
	//base.Logger.Debugf("回复消息:", string(msg))
	return true
}

//错误消息
func WS_SendError(code int, msg string, ws *websocket.Conn) bool {

	//错误
	type TagServer_Error struct {
		MessageId int    `json:"messageId"` //消息ID
		ErrorCode int    `json:"errorCode"` //错误码
		ErrorMsg  string `json:"errorMsg"`  //错误信息
	}

	WSdata := &TagServer_Error{
		SERVER_MSG_ID_ERROR,
		code,
		msg,
	}
	return WS_SendMessage(WSdata, ws)
}
