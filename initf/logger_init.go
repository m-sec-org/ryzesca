package initf

import (
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// LoggerInit  日志初始化
func LoggerInit() *zap.SugaredLogger {
	logMode := zapcore.ErrorLevel
	if viper.GetBool("development.develop") {
		logMode = zapcore.DebugLevel
	}
	core := zapcore.NewCore(
		getEncoder(),
		zapcore.NewMultiWriteSyncer(getLogWriter()),
		logMode)
	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
func getLogWriter() zapcore.WriteSyncer {
	stSeparator := string(os.PathSeparator)
	stRootDir, _ := os.Getwd()
	filePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly) + ".log"
	// 分割器
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    viper.GetInt("log.MaxSize"),    // 最大M数，超过则分割
		MaxBackups: viper.GetInt("log.MaxBackups"), // 最大备份数量
		MaxAge:     viper.GetInt("log.MaxAge"),     // 最大保留天数
		Compress:   false,                          // 是否压缩 disabled by default
	}
	return zapcore.AddSync(lumberJackLogger)
}
