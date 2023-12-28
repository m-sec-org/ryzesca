package main

import (
	"RyzeSCA/global"
	"RyzeSCA/initf"
)

func main() {

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
	initf.Banner()
	initf.GrpcServerInit()
}
