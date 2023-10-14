package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

/*
	将日志输出至文件中
 */

func main2() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(config)

	tmpLog, _ := os.Create("tmp.log")
	writeSyncer := zapcore.AddSync(tmpLog)

	core := zapcore.NewCore(encoder, writeSyncer, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Debug("CustErr", zap.String("err", "errMsg"))
}
