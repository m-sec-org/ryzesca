package initf

import (
	"fmt"
	"github.com/spf13/viper"
)

// ViperInit 配置信息初始化
func ViperInit() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		//global.Logger.Errorf("配置文件读取失败" + err.Error())
		fmt.Printf("配置文件读取失败" + err.Error())
		panic("配置文件读取失败" + err.Error())
	}
}
