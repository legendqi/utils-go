/* coding: utf-8
@Time :   2021/3/4 下午2:42
@Author : legend
@File :   log.go
*/
package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

var (
	logLevel    = zapcore.InfoLevel
	logHandlers = []string{"file", "console"}
	Logger      *zap.SugaredLogger
)

func init() {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	consoleLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})
	// 实现两个判断日志等级的interface
	runningLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == logLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel
	})
	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	var zapcores []zapcore.Core
	for _, handler := range logHandlers {
		if handler == "console" {
			hook := os.Stdout
			zapcores = append(
				zapcores,
				zapcore.NewCore(encoder, zapcore.AddSync(hook), consoleLevel),
			)
		} else if handler == "file" {
			runningHook := getWriter("./logs/running.log")
			errorHook := getWriter("./logs/error.log")
			zapcores = append(
				zapcores,
				zapcore.NewCore(encoder, zapcore.AddSync(runningHook), runningLevel),
				zapcore.NewCore(encoder, zapcore.AddSync(errorHook), errorLevel),
			)
		}
	}
	// 最后创建具体的Logger
	core := zapcore.NewTee(zapcores...)
	// 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	Logger = logger.Sugar()
	defer logger.Sync()
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 info.log.YYmmddHH
	// info.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y-%m-%d", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

func DebugFormatter(template string, args ...interface{}) {
	Logger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func InfoFormatter(template string, args ...interface{}) {
	Logger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

func WarnFormatter(template string, args ...interface{}) {
	Logger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}

func ErrorFormatter(template string, args ...interface{}) {
	Logger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	Logger.DPanic(args...)
}

func DPanicFormatter(template string, args ...interface{}) {
	Logger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	Logger.Panic(args...)
}

func PanicFormatter(template string, args ...interface{}) {
	Logger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

func FatalFormatter(template string, args ...interface{}) {
	Logger.Fatalf(template, args...)
}

//设置日志等级和handler
func SetLogConfig(level string, handlers []string) {
	formatterLevel := strings.ToLower(level)
	switch formatterLevel {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	case "panic":
		logLevel = zapcore.PanicLevel
	case "fatal":
		logLevel = zapcore.FatalLevel
	default:
		logLevel = zapcore.InfoLevel
	}
	logHandlers = handlers
}

func SetLevel(level string) {
	formatterLevel := strings.ToLower(level)
	switch formatterLevel {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	case "panic":
		logLevel = zapcore.PanicLevel
	case "fatal":
		logLevel = zapcore.FatalLevel
	default:
		logLevel = zapcore.InfoLevel
	}
}
