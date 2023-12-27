package initf

import (
	"RyzeSCA/global"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// MysqlInit  mysql初始化
func MysqlInit() (*gorm.DB, error) {
	logModel := logger.Info
	if viper.GetBool("development.develop") {
		logModel = logger.Warn
	}
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "sys_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logModel),
	})
	if err != nil {
		global.Logger.Error("MySQL连接失败", err)
		return nil, err
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(viper.GetInt("mysql.MaxIdleCons"))
	sqlDb.SetMaxOpenConns(viper.GetInt("mysql.MaxOpenCons"))
	sqlDb.SetConnMaxLifetime(time.Minute * 60)

	if err != nil {
		return nil, err
	}
	return db, err
}
