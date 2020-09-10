package myWebscoket

import (
	"errors"
	"github.com/HappyTeemo7569/mymod/base"
	"golang.org/x/net/websocket"
	"time"
)

//获取连接信息
func GetSocketItem(ws *websocket.Conn) *SocketItem {
	var item *SocketItem
	value, _ := SocketList.Load(ws)
	if value != nil {
		item = value.(*SocketItem)
	}
	return item
}

//添加连接
func AddSocketItem(ws *websocket.Conn) {
	SocketList.Store(ws, &SocketItem{
		true,
		nil,
		int(time.Now().Unix()),
	})
}

//删除连接
func DelSocketItem(ws *websocket.Conn) {
	SocketList.Delete(ws)
}

//检测socket连接
func CheckSocketList() {
	timenow := int(time.Now().Unix())

	SocketList.Range(func(k, v interface{}) bool {
		switch v.(type) {
		case *SocketItem:
			value := v.(*SocketItem)
			if value != nil && (!value.BConnect && timenow-value.MsgTime >= 12) {
				value.BConnect = false
				base.Logger.Infof("超时断开用户，当前时间:%d，用户时间:%d", timenow, value.MsgTime)
				ws := k.(*websocket.Conn)
				NormalChan <- WsMsgItem{CONN_TYPE_CLOSE, ws, nil}
				return true
			}
		}
		return false
	})
}

//设置用户数据
func SetSocketItem(ws *websocket.Conn, userData interface{}) (bool, error) {
	skt := GetSocketItem(ws)
	if skt == nil {
		return false, errors.New("未找到该连接")
	}
	if skt.UserData != nil {
		return false, errors.New("该连接已有数据")
	}
	skt.UserData = userData
	return true, nil
}

//更新用户数据
func UpdateSocketItem(ws *websocket.Conn, userData interface{}) (bool, error) {
	skt := GetSocketItem(ws)
	if skt == nil {
		return false, errors.New("未找到该连接")
	}
	if skt.UserData == nil {
		return false, errors.New("该连接还未设置数据")
	}
	skt.UserData = userData
	return true, nil
}

//获取用户数据
func GetUserData(ws *websocket.Conn, userData interface{}) (interface{}, error) {
	skt := GetSocketItem(ws)
	if skt == nil {
		return nil, errors.New("未找到该连接")
	}
	if skt.UserData == nil {
		return nil, errors.New("该连接还未设置数据")
	}
	return skt.UserData, nil
}
