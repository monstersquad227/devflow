package middleware

/**
import (
	"bytes"
	"devflow/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"
)

type LogType string

const (
	HttpIn LogType = "HttpIn" // Http请求入口日志
	Panic  LogType = "Panic"  // Panic错误日志
)

// var logger = logrus.New()
var logger *logrus.Logger

// CustomJSONFormatter 自定义 JSON 格式化器
type CustomJSONFormatter struct{}

func (f *CustomJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	logEntry := make(map[string]interface{})

	// 仅保留有意义的字段，避免 logrus 自动添加 msg/fields.level
	for key, value := range entry.Data {
		logEntry[key] = value
	}

	// 手动添加时间
	logEntry["time"] = entry.Time.Format(time.RFC3339)

	// 序列化为 JSON
	logBytes, err := json.Marshal(logEntry)
	if err != nil {
		return nil, err
	}
	return append(logBytes, '\n'), nil
}

// setupLogger 初始化日志
func setupLogger() {
	// 确保日志目录存在
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			panic(fmt.Sprintf("创建日志目录失败: %v", err))
		}
	}
	// 生成完整的日志文件路径
	//fmt.Println(config.GlobalConfig.Application.Name + "applicationName")
	logFileName := filepath.Join(logDir, fmt.Sprintf("%s_%s.log",
		config.GlobalConfig.Application.Name,
		time.Now().Format("2006-01-02"),
	))

	// 配置 lumberjack
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    100, // MB
		MaxBackups: 30,  // 保留旧日志文件数量
		MaxAge:     30,  // 保留天数
		Compress:   true,
		LocalTime:  true,
	}
	logger = logrus.New()
	logger.SetFormatter(&CustomJSONFormatter{})
	logger.SetOutput(lumberjackLogger)
	logger.SetLevel(logrus.InfoLevel)
	// 测试日志输出
	//logger.WithFields(logrus.Fields{
	//	"event": "logger_initialized",
	//}).Info("Logger 初始化成功")

}

// Logger 记录 HTTP 请求日志
func Logger() gin.HandlerFunc {
	setupLogger()

	return func(c *gin.Context) {
		start := time.Now()

		// 读取请求体
		var requestBody map[string]interface{}
		if c.Request.Body != nil {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				_ = json.Unmarshal(bodyBytes, &requestBody)
				// 重新赋值给 c.Request.Body，确保后续可以再次读取
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		c.Next()
		fmt.Println("[自定义中间件] 正在记录请求日志...")

		// 计算延迟时间（ms）
		latencyTime := time.Since(start).Milliseconds()

		requestLog := map[string]interface{}{
			"logType":              HttpIn,
			"context":              "devflow",
			"requestMethod":        c.Request.Method,
			"requestUri":           c.Request.RequestURI,
			"remoteAddr":           c.ClientIP(),
			"requestContentLength": c.Request.ContentLength,
			"userAgent":            c.Request.UserAgent(),
			"requestHeaders":       c.Request.Header,
			"requestBody":          requestBody,
			"requestParameters":    c.Request.URL.Query(),
			"responseStatus":       c.Writer.Status(),
			"responseTime":         latencyTime,
		}

		logger.WithFields(requestLog).Info()
	}
}

// RecoveryWithLogger 捕获 `panic` 并写入日志
func RecoveryWithLogger() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		stackTrace := string(debug.Stack())
		panicLog := map[string]interface{}{
			"logType":    Panic,
			"stackTrace": stackTrace,
			"requestUri": c.Request.RequestURI,
			"remoteAddr": c.ClientIP(),
		}
		logger.WithFields(panicLog).Error()
		c.AbortWithStatus(500)
	})
}

**/
