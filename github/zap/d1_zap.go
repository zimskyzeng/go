package main

import (
	"go.uber.org/zap"
)

/*
	zap是高性能、支持结构化数据、可分日志级别的包
	如果需要使用一些个性化的配置，比方说将日志写入到文件中、切割日志文件大小 需要用到go.uber.org/zap/zapcore
 */

func main(){
	logger, _ := zap.NewProduction()
	defer logger.Sync()  // 用来刷新缓存，一定要有这个

	// 用法1 生成Sugar来使用
	sugar := logger.Sugar()
	// 以给定格式的方式输出
	sugar.Infof("err:%s, msg: %s", "syncErr", "errMsg")
	// 以key-value方式输出的结构化日志，第一位是日志名称
	sugar.Infow("syncErr", "err", "errMsg")

	// 用法2 直接使用Logger对象 性能会优于Sugar，但是日志的数据必须是指定类型的
	logger.Info("SyncErr",
		zap.String("stringMsg", "outputStringMsg"),
		zap.Int("intMsg", 502),
	)
}
