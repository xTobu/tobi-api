package log

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.Logger
	atom   zap.AtomicLevel
)

func init() {

	// 設置日誌級別
	SetLogLevel("info")

	// 初始化 zap logger
	logger = initLogger()

	/* Using Demo */
	////////////////////
	// logger.Info("無法獲取網址",
	//  zap.Error(err)
	// 	zap.String("url", "http://www.baidu.com"),
	// 	zap.Int("attempt", 3),
	// 	zap.Duration("backoff", time.Second))
}

func initLogger() *zap.Logger {

	fileName := fmt.Sprintf("./logs/%04d-%02d-%02d.log", time.Now().Year(), time.Now().Month(), time.Now().Day())
	hook := lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    256,  // megabytes
		MaxBackups: 3,    // 最多保存多少個備份
		MaxAge:     28,   // days
		Compress:   true, // 是否壓縮
		LocalTime:  true,
	}
	syncWriter := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout), // for console
		zapcore.AddSync(&hook),     // for lumberjack.logger
	)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 大寫寫編碼器
		EncodeTime:     zapcore.ISO8601TimeEncoder,  // ISO8601 UTC 時間格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 簡短路徑的編碼器
		EncodeName:     zapcore.FullNameEncoder,
	}

	//設置 zap 核心
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // 輸出格式配置 Console
		// zapcore.NewJSONEncoder(encoderConfig), // 輸出格式配置 JSON
		syncWriter,
		atom, // 日誌級別
	)

	// 開啟開發模式，堆棧跟踪
	caller := zap.AddCaller()

	// AddCallerSkip(1) : 因為呼叫上有一層封裝，故 Caller 必須往上跳一層
	callerSkip := zap.AddCallerSkip(1)

	// 開啟文件及行號
	development := zap.Development()

	// 設置初始化字段
	// filed := zap.Fields(zap.String("serviceName", "serviceName"))

	// 建置日誌
	logger = zap.New(core, development, caller, callerSkip)

	return logger
}

// SetLogLevel : set log level
func SetLogLevel(logLevel string) {

	atom = zap.NewAtomicLevel()

	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "info":
	default:
		level = zap.InfoLevel
	}

	atom.SetLevel(level)
}

// Debug : just pack a function, named from zap
func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

// Info : just pack a function, named from zap
func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

// Warn : just pack a function, named from zap
func Warn(msg string, fields ...zapcore.Field) {
	logger.Warn(msg, fields...)
}

// Error : just pack a function, named from zap
func Error(msg string, err error) {
	logger.Error(msg, zap.Error(err))
}

// DPanic : just pack a function, named from zap
func DPanic(msg string, fields ...zapcore.Field) {
	logger.DPanic(msg, fields...)
}

// Panic : just pack a function, named from zap
func Panic(msg string, fields ...zapcore.Field) {
	logger.Panic(msg, fields...)
}

// Fatal : just pack a function, named from zap
func Fatal(msg string, fields ...zapcore.Field) {
	logger.Fatal(msg, fields...)
}
