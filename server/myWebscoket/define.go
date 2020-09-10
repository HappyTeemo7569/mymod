package myWebscoket

import (
	"github.com/HappyTeemo7569/mymod/server/define"
	"golang.org/x/net/websocket"
	"sync"
)

//网络连接
var SocketList sync.Map

const (
	HEARTBEAT_MSG_ID    = 100 //心跳包
	SERVER_MSG_ID_ERROR = 200 //错误消息
)

//wsSocket连接体
type SocketItem struct {
	BConnect bool        //是否连接
	UserData interface{} //用户数据
	MsgTime  int         //时间戳，重连用
}

var NormalChan QueChan

//消息通道
type QueChan chan WsMsgItem

//消息体
type WsMsgItem struct {
	Conntype int
	Ws       *websocket.Conn
	WSdata   *TagClient_Universal //客户端消息
}

//客户端消息通用结构
type TagClient_Universal struct {
	InterfaceNumber int                  `json:"interfaceNumber"` //请求消息ID
	Data            define.ClientMessage `json:"data"`
}

//消息体类型
const (
	CONN_TYPE_CONNECT = iota //连接
	CONN_TYPE_READ           //业务逻辑
	CONN_TYPE_CLOSE          //关闭
	CONN_TYPE_CHECK          //校验
)

//心跳包服务端回包
type tagServer_HeartBeat struct {
	MessageId int `json:"messageId"` //消息ID
}
