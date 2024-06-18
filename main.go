package main

import (
	"os"

	"github.com/calendarproject/common"
	"github.com/calendarproject/router"
	"github.com/calendarproject/ws"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	//初始化配置
	InitConfig()

	//初始化数据库
	common.InitDB()

	//使用gin
	r := gin.Default()
	r = router.CollectRoute(r)
	// 注册WebSocket路由
	r.GET("/ws", ws.WebSocketHandler)
	port := viper.GetString("server.port")

	if port != "" {
		_ = r.Run(":" + port)
	} else {
		err := r.Run()
		if err != nil {
			return
		}
	}
}

func InitConfig() {

	workDir, _ := os.Getwd()

	viper.SetConfigName("application")

	viper.SetConfigType("yml")

	viper.AddConfigPath(workDir + "/config")

	err := viper.ReadInConfig()

	if err != nil {

		return

	}

}
