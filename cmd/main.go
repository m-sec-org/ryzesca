package main

import (
	"RyzeSCA/global"
	"RyzeSCA/initf"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func main() {
	fmt.Println(time.Now(), "grpc 启动成功 端口号为", viper.GetInt("server.port"))
}

func init() {
	initf.ViperInit()
	logger := initf.LoggerInit()
	global.Logger = logger
	mysql, err := initf.MysqlInit()
	if err != nil {
		global.Logger.Error("mysql初始化失败", err)
		return
	}
	global.MysqlDB = mysql
	global.Logger.Info("mysql连接成功")
	fmt.Println(time.Now().Format("2006-01-02 :15-04-05"), "grpc 启动成功 端口号为", viper.GetInt("server.port"))
	initf.GrpcServerInit()
}
