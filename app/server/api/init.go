package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/ismdeep/go-rest-api-demo/app/server/conf"
	"github.com/ismdeep/go-rest-api-demo/pkg/log"
	"github.com/ismdeep/go-rest-api-demo/pkg/quantumid"
)

var eng *gin.Engine

func init() {
	gin.SetMode(conf.Basic.Server.Mode)
	eng = gin.Default()
	eng.Use(traceLoggerMiddleware())

	// auth
	eng.GET("/api/v1/auth/profile", MyProfile)
	eng.POST("/api/v1/auth/sign-up", SignUp)
	eng.POST("/api/v1/auth/sign-in", SignIn)

	// users
	eng.GET("/api/v1/users", GetUserList)
}

// Run http server
func Run() {
	if err := eng.Run(conf.Basic.Server.Bind); err != nil {
		panic(err)
	}
}

func traceLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 每个请求生成的请求traceId具有全局唯一性
		log.AddTraceID(ctx, quantumid.NewString())

		log.WithContext(ctx).Info("init",
			zap.String("method", ctx.Request.Method),
			zap.String("url", ctx.Request.URL.String()),
		)

		ctx.Next()
	}
}
