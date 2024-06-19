package ws

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 客户端连接
type Client struct {
	conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	// 设置 Upgrader 的参数,如缓冲区大小、子协议等
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// 客户端连接列表
var clients = make(map[uint64]*Client)

var unSendMsg = make(map[uint64]map[string]bool)

var mutex = sync.Mutex{}

func WebSocketHandler(c *gin.Context) {
	// 获取WebSocket连接
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	// 创建新客户端连接
	client := &Client{
		conn: ws,
	}

	// 从url `...ws?uid=1`获取uid
	uId, err := strconv.ParseUint(c.Query("uid"), 10, 64)

	fmt.Println(uId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println("uid获取异常", err)
	}
	mesType := c.Query("mesType")

	// 添加客户端连接到列表
	addClient(uId, client)
	mes := unSendMsg[uId]
	if len(mes) != 0 {
		for key, _ := range mes {
			SendMessage(uId, key, 1)
		}
	}
	// 处理WebSocket消息
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			// 从列表中删除客户端连接
			removeClient(uId)
			break
		}

		fmt.Println("messageType:", messageType)
		fmt.Println("p:", string(p))

		if mesType == "all" {
			// 向客户端推送消息
			broadcastMessage(p)
		} else {
			//自己发给自己
			SendMessage(uId, string(p), 0)
		}

	}

	// 关闭WebSocket连接
	ws.Close()
}

func addClient(key uint64, client *Client) {
	mutex.Lock()
	defer mutex.Unlock()
	clients[key] = client
}

func addUnSend(key uint64, message string) {
	mutex.Lock()
	defer mutex.Unlock()

	// 检查 key 是否存在于 unSendMessage 中
	messages, ok := unSendMsg[key]
	if !ok {
		// 如果 key 不存在,创建一个新的 map 并添加 message
		messages = make(map[string]bool)
		messages[message] = true
		unSendMsg[key] = messages
	} else {
		// 如果 key 已经存在,检查 message 是否已经存在
		if _, exists := messages[message]; !exists {
			// 如果 message 不存在,添加到 messages 中
			messages[message] = true
		}
	}
}

func removeUnSend(key uint64, mes string) {
	mutex.Lock()
	defer mutex.Unlock()
	// 检查 key 是否存在于 unSendMessage 中
	messages, ok := unSendMsg[key]
	if ok {
		delete(messages, mes)
	}
}

func removeClient(key uint64) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(clients, key)
}

// 所有人推送
func broadcastMessage(message []byte) {
	mutex.Lock()
	defer mutex.Unlock()
	for key, client := range clients {
		err := client.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			removeClient(key)
		}
	}
}

// 向某个人发送信息 0第一次发 1重发
func SendMessage(key uint64, message string, sendType uint) {
	mutex.Lock()
	defer mutex.Unlock()
	client := clients[key]
	err := client.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		//保存未发送的信息
		addUnSend(key, message)
		removeClient(key)
	}
	if sendType == 1 {
		removeUnSend(key, message)
	}
}
