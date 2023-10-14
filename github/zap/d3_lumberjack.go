package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
	将日志输出到文件中并且自动切割
*/

func main3(){
	encoder := GetEncoder()
	writeSync := GetWriteSync()
	core := zapcore.NewCore(encoder, writeSync, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())  // 添加显示调用处
	defer logger.Sync()

	logger.Error("CustErr", zap.String("err", "errors"))
}

// 创建Encoder
func GetEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder  // 优化时间显示
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(config)
	return encoder
}

func GetWriteSync () zapcore.WriteSyncer{
	logger := &lumberjack.Logger{
		Filename:   "tmp.log",
		MaxSize:    10,
		MaxAge:     1,
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   false,
	}
	return zapcore.AddSync(logger)
}