package log

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TraceIDKeyType trace id key type
type TraceIDKeyType string

const (
	// TraceIDKey trace id key
	TraceIDKey TraceIDKeyType = "traceId"
)

// Logger instance
var Logger *zap.Logger

// 初始化日志配置
func init() {
	Logger = zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
			zap.DebugLevel,
		),
	)
}

// AddTraceID 给ctx增加 traceID
func AddTraceID(ctx *gin.Context, traceID string) {
	ctx.Set(string(TraceIDKey), traceID)
}

// WithContext 从指定的context返回一个zap实例（关键方法）
func WithContext(ctx context.Context) *zap.Logger {
	if v := ctx.Value(TraceIDKey); v != "" {
		if s, ok := v.(string); ok {
			return Logger.With(zap.String("traceId", s))
		}
	}

	if v := ctx.Value(string(TraceIDKey)); v != "" {
		if s, ok := v.(string); ok {
			return Logger.With(zap.String("traceId", s))
		}
	}

	return Logger
}

// NewTraceContext new context with a traceID
func NewTraceContext(traceID string) context.Context {
	return context.WithValue(context.Background(), TraceIDKey, traceID)
}
