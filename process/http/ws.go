package http

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"zrw/goMin/process/controller"
)

// Ws webSocket请求ping 返回pong
func Ws(ctx *gin.Context) {
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		if string(message) == "server_time" {
			resp, _, _ := controller.GetServerTime()
			serverTime := strconv.FormatInt(resp.ServerTime, 10)
			message = []byte(serverTime)
		}
		//写入数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
