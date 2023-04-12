// @User CPR
package utils

import (
	"VideoEdit/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

var (
	Logger *zap.Logger
)

func InitLogger() {
	// 生成日志文件目录
	_ = CreateDir(config.Cfg.Zap.Directory)
	core := zapcore.NewCore(getEncoder(), getWriterSyncer(), getLevelPriority())
	Logger = zap.New(core)

	if config.Cfg.Zap.ShowLine {
		// 获取 调用的文件, 函数名称, 行号
		Logger = Logger.WithOptions(zap.AddCaller())
	}

	log.Println("Zap Logger 初始化成功")
}

// 编码器: 如何写入日志
func getEncoder() zapcore.Encoder {
	// 参考: zap.NewProductionEncoderConfig()
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",                      // 消息的key
		LevelKey:       "level",                        // 级别的key
		TimeKey:        "time",                         // 时间的key
		NameKey:        "logger",                       // 日志器的key
		CallerKey:      "caller",                       // 调用者的key
		StacktraceKey:  "stacktrace",                   // 堆栈跟踪的key
		LineEnding:     zapcore.DefaultLineEnding,      // 换行符
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 大写编码器
		EncodeTime:     customTimeEncoder,              // 自定义时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, // 秒级别编码器
		EncodeCaller:   zapcore.FullCallerEncoder,      // ?	// 全路径编码器
	}

	if config.Cfg.Zap.Format == "json" {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 日志输出路径: 文件、控制台、双向输出
func getWriterSyncer() zapcore.WriteSyncer {
	file, err := os.Create(config.Cfg.Zap.Directory + "/log_" + time.Now().Format("2006-01-02") + ".log")
	if err != nil {
		panic(err)
	}
	// 双向输出
	if config.Cfg.Zap.LogInConsole {
		fileWriter := zapcore.AddSync(file)
		consoleWriter := zapcore.AddSync(os.Stdout)
		return zapcore.NewMultiWriteSyncer(fileWriter, consoleWriter)
	}

	// 输出到文件
	return zapcore.AddSync(file)
}

// 获取日志输出级别
func getLevelPriority() zapcore.LevelEnabler {
	switch config.Cfg.Zap.Level {
	case "debug", "Debug":
		return zap.DebugLevel
	case "info", "Info":
		return zap.InfoLevel
	case "warn", "Warn":
		return zap.WarnLevel
	case "error", "Error":
		return zap.ErrorLevel
	case "dpanic", "DPanic":
		return zap.DPanicLevel
	case "panic", "Panic":
		return zap.PanicLevel
	case "fatal", "Fatal":
		return zap.FatalLevel
	}
	return zap.InfoLevel
}

// 自定义日志输出时间格式
func customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(config.Cfg.Zap.Prefix + t.Format("2006/01/02 - 15:04:05"))
}
