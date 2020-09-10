package define

import "golang.org/x/net/websocket"

//关闭通道
var CloseChan chan int

//客户端业务逻辑消息
type ClientMessage interface{}

//业务逻辑队列
var JoinChan JoinQueChan

type JoinQueChan chan JoinMsgItem
type JoinMsgItem struct {
	Ws     *websocket.Conn
	WSdata interface{}
}

//退出业务逻辑队列（关闭连接会写）
var ExitChan ExitGameQueChan

type ExitGameQueChan chan ExitGameMsgItem
type ExitGameMsgItem struct {
	UserData interface{} //会发送存储的用户数据
}
