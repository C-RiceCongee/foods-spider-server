package pkg

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

// log 输出编码配置！
func setEncoder() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if Config.Mode == "debug" {
		// 	开发模式下使用 console 格式 看起来更友好！
		//	CapitalColorLevelEncoder 带颜色的Level
		encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encodeConfig)
	} else {
		// 	生产模式就使用JSON 格式！
		encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		return zapcore.NewJSONEncoder(encodeConfig)
	}
}

/*
setIoWriter

	stdout   是否需要标准输出到控制台
	filename 输出到的文件路径
	TODO feature 全部参数化！
*/
func setIoWriter(stdout bool, filename string) zapcore.WriteSyncer {
	lg := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    Config.LogConfig.MaxSize,
		MaxBackups: Config.LogConfig.MaxBackups,
		MaxAge:     Config.LogConfig.MaxAge,
		Compress:   Config.LogConfig.Compress,
	}
	var ws io.Writer
	if stdout {
		ws = io.MultiWriter(lg, os.Stdout)
	} else {
		ws = lg
	}
	return zapcore.AddSync(ws)
}

func InitLogger() (logger *zap.Logger, err error) {
	/*
		DebugLevel 等级的日志输出包括了 ERROR 等级, 同样可以把ERROR的同时输出到控制台和默认文件
		所以下面的ErrorLevel 只需要输出到文件中就可以了！
	*/

	// 读取配置基础日志级别,支持将 字符串 转换为 zap.Level 类型~
	level := new(zapcore.Level)
	err = level.UnmarshalText([]byte(Config.LogConfig.Level))
	if err != nil {
		return nil, err
	}
	coreDebug := zapcore.NewCore(setEncoder(), setIoWriter(true, Config.LogConfig.Filename), level)
	coreError := zapcore.NewCore(setEncoder(), setIoWriter(false, Config.LogConfig.ErrorFilename), zap.ErrorLevel)
	core := zapcore.NewTee(coreDebug, coreError)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(Config.LogConfig.CallerSkip))
	return logger, nil
}
