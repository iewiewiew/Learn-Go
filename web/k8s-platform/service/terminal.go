package service

import (
	"github.com/gorilla/websocket"
	"k8s.io/client-go/tools/remotecommand"
	"net/http"
	"time"
)

/**
 * @author       weimenghua
 * @time         2024-07-29 15:48
 * @description
 */

// 终止符
const END_OF_TRANSMISSION = "\u0004"

type TerminalMessage struct {
	Operation string `json:"operation"`
	Data      string `json:"data"`
	Rows      uint16 `json:"rows"`
	Cols      uint16 `json:"cols"`
}

// 初始化一个websocket.Upgrader类型的对象，用于http协议升级为websocket协议
var upgrader = func() websocket.Upgrader { // 定义了一个匿名函数,并将其赋值给变量 upgrader。这个匿名函数返回一个 websocket.Upgrader 类型的值。
	upgrader := websocket.Upgrader{} // 创建了一个新的 websocket.Upgrader 值,并将其赋值给变量 upgrader。这个变量与第一行定义的 upgrader 变量是同一个。
	upgrader.HandshakeTimeout = time.Second * 2
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	return upgrader
}

type TerminalMessage2 struct {
	wsConn   *websocket.Conn
	sizeChan chan remotecommand.TerminalSize
	doneChan chan struct{}
	tty      bool
}
