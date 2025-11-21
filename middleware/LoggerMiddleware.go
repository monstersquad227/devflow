package middleware

import (
	"bytes"
	"devflow/config"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"
	"sync"
	"time"
)

type LogType string

const (
	HttpIn  LogType = "HttpIn"  // Http请求入口日志
	HttpOut LogType = "HttpOut" // Http请求出口日志（调用外部服务）
	Panic   LogType = "Panic"   // Panic错误日志
)

var (
	logger       *logrus.Logger
	loggerMutex  sync.RWMutex
	lastLogDate  string
	rotateWriter io.WriteCloser
)

// CustomJSONFormatter 自定义 JSON 格式化器
type CustomJSONFormatter struct{}

func (f *CustomJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	logEntry := make(map[string]interface{})

	// 保留所有字段
	for key, value := range entry.Data {
		logEntry[key] = value
	}

	// 添加时间和级别
	logEntry["time"] = entry.Time.Format(time.RFC3339)
	logEntry["level"] = entry.Level.String()

	// 序列化为 JSON
	logBytes, err := json.Marshal(logEntry)
	if err != nil {
		return nil, err
	}
	return append(logBytes, '\n'), nil
}

// initLogger 初始化 logger（支持日期变更时重新初始化）
func initLogger() {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()

	currentDate := time.Now().Format("20060102")

	// 如果 logger 已存在且日期未变，直接返回
	if logger != nil && lastLogDate == currentDate {
		return
	}

	// 如果日期变了，关闭旧的 writer
	if rotateWriter != nil {
		rotateWriter.Close()
	}

	// 确保日志目录存在
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(fmt.Sprintf("创建日志目录失败: %v", err))
	}

	// 使用通配符模式，让 rotatelogs 自动管理文件名
	logFileName := filepath.Join(logDir, fmt.Sprintf("%s_%%Y%%m%%d.log",
		config.GlobalConfig.Application.Name,
	))

	// 正确的轮转配置
	var err error
	rotateWriter, err = rotatelogs.New(
		logFileName,
		rotatelogs.WithClock(rotatelogs.Local),    // 使用本地时间
		rotatelogs.WithRotationTime(24*time.Hour), // 每天轮转
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 保留 7 天
	)
	if err != nil {
		panic(fmt.Sprintf("无法初始化日志轮转: %v", err))
	}

	// 创建 logger
	logger = logrus.New()
	logger.SetFormatter(&CustomJSONFormatter{})

	// 同时输出到控制台和文件
	multiWriter := io.MultiWriter(os.Stdout, rotateWriter)
	logger.SetOutput(multiWriter)
	logger.SetLevel(logrus.InfoLevel)

	// 记录当前日期
	lastLogDate = currentDate
}

// Logger 记录 HTTP 请求日志
func Logger(logType LogType) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 每次请求都检查是否需要重新初始化
		initLogger()

		start := time.Now()

		// 读取请求体
		var requestBody interface{}
		if c.Request.Body != nil && c.Request.ContentLength > 0 {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				// 检查 unmarshal 错误
				if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
					requestBody = string(bodyBytes) // 如果不是 JSON，保存原始字符串
				}
				// 重新赋值给 c.Request.Body
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		// 执行下一个中间件/处理器
		c.Next()

		// 计算延迟时间（ms）
		latencyTime := time.Since(start).Milliseconds()

		// 只在 HttpIn 类型时记录
		if logType == HttpIn {
			loggerMutex.RLock()
			defer loggerMutex.RUnlock()

			if logger != nil {
				httpInEntry := logrus.Fields{
					"logType":              string(HttpIn),
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
				logger.WithFields(httpInEntry).Info("HTTP Request")
			}
		}
	}
}

// LogHttpOut 记录向外部发起的 HTTP 请求
func LogHttpOut(
	method string,
	url string,
	requestHeaders http.Header,
	requestBody interface{},
	responseStatus int,
	responseBody interface{},
	latency time.Duration,
	err error) {
	initLogger()

	loggerMutex.RLock()
	defer loggerMutex.RUnlock()

	if logger == nil {
		return
	}

	httpOutEntry := logrus.Fields{
		"logType":        string(HttpOut),
		"context":        "devflow",
		"requestMethod":  method,
		"requestUrl":     url,
		"requestHeaders": requestHeaders,
		"requestBody":    requestBody,
		"responseStatus": responseStatus,
		"responseBody":   responseBody,
		"responseTime":   latency.Milliseconds(),
	}

	if err != nil {
		httpOutEntry["error"] = err.Error()
		logger.WithFields(httpOutEntry).Error("HTTP Request Failed")
	} else {
		logger.WithFields(httpOutEntry).Info("HTTP Request")
	}
}

// RecoveryWithLogger 捕获 `panic` 并写入日志
func RecoveryWithLogger() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		// 确保 logger 已初始化
		initLogger()

		loggerMutex.RLock()
		defer loggerMutex.RUnlock()

		if logger == nil {
			c.JSON(500, gin.H{
				"code":    500,
				"message": "Internal Server Error",
			})
			return
		}

		stackTrace := string(debug.Stack())

		panicLog := logrus.Fields{
			"logType":    string(Panic),
			"error":      fmt.Sprintf("%v", err),
			"stackTrace": stackTrace,
			"requestUri": c.Request.RequestURI,
			"remoteAddr": c.ClientIP(),
			"method":     c.Request.Method,
		}

		logger.WithFields(panicLog).Error("Panic Recovered")

		// 返回 JSON 错误响应
		c.JSON(500, gin.H{
			"code":    500,
			"message": "Internal Server Error",
		})
	})
}
