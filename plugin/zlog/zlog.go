package zlog

import (
	"context"
	"github.com/ct-llper/go-pkg/plugin/metadata"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

// error logger
var zapLogger *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func Init(fileName string, levelString string) {
	file, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(file)
	//暂时将所有输出重定向到该文件
	//os.Stdout = file
	//os.Stderr = file
	level := getLoggerLevel(levelString)
	syncWriter := zapcore.AddSync(file)
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zapLogger = logger.Sugar()
}

func Debug(ctx context.Context, args ...interface{}) {
	withs := GetLogId(ctx)
	zapLogger.With(withs[:]...).Debug(args...)
}

func Debugf(ctx context.Context, template string, args ...interface{}) {
	withs := GetLogId(ctx)
	zapLogger.With(withs[:]...).Debugf(template, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	withs := GetLogId(ctx)
	zapLogger.With(withs[:]...).Info(args...)
}

func Infof(ctx context.Context, template string, args ...interface{}) {
	withs := GetLogId(ctx)
	zapLogger.With(withs[:]...).Infof(template, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	withs := GetLogId(ctx)
	zapLogger.With(withs[:]...).Warn(args...)
}

func Warnf(ctx context.Context, template string, args ...interface{}) {
	withs := GetLogId(ctx)
	zapLogger.With(withs[:]...).Warnf(template, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	withs := GetLogId(ctx)
	zapLogger.With(withs[:]...).Error(args...)
}

func Errorf(ctx context.Context, template string, args ...interface{}) {
	withs := GetLogId(ctx)
	zapLogger.With(withs[:]...).Errorf(template, args...)
}

func Panic(ctx context.Context, args ...interface{}) {
	withs := GetLogId(ctx)
	zapLogger.With(withs[:]...).Panic(args...)
}

func Panicf(ctx context.Context, template string, args ...interface{}) {
	withs := GetLogId(ctx)
	zapLogger.With(withs[:]...).Panicf(template, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	withs := GetLogId(ctx)
	zapLogger.With(withs[:]...).Fatal(args...)
}

func Fatalf(ctx context.Context, template string, args ...interface{}) {
	withs := GetLogId(ctx)
	zapLogger.With(withs[:]...).Fatalf(template, args...)
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func GetLogId(ctx context.Context) []interface{} {
	withs := make([]interface{}, 2)
	if ctx != nil {
		md, ok := metadata.FromContext(ctx)
		if ok {
			withs[0] = "request-id"
			withs[1] = md["Request-Id"]
		} else {
			withs[0] = "request-id"
			withs[1] = "Not Found request-id"
		}
	} else {
		withs[0] = "request-id"
		withs[1] = "xxx"
	}
	return withs
}
