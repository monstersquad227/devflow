package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
	"time"
)

var (
	// HTTP 请求总数
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	// HTTP 请求延迟
	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latencies in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint", "status"},
	)

	// 当前正在处理的请求数
	httpRequestsInFlight = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Current number of HTTP requests being processed",
		},
	)

	// HTTP 请求大小
	httpRequestSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_size_bytes",
			Help:    "HTTP request sizes in bytes",
			Buckets: prometheus.ExponentialBuckets(100, 10, 7),
		},
		[]string{"method", "endpoint"},
	)

	// HTTP 响应大小
	httpResponseSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_size_bytes",
			Help:    "HTTP response sizes in bytes",
			Buckets: prometheus.ExponentialBuckets(100, 10, 7),
		},
		[]string{"method", "endpoint"},
	)
)

// PrometheusMonitor Prometheus 监控中间件
func PrometheusMonitor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 增加正在处理的请求数
		httpRequestsInFlight.Inc()
		defer httpRequestsInFlight.Dec()

		start := time.Now()
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}
		method := c.Request.Method

		// 记录请求大小
		if c.Request.ContentLength > 0 {
			httpRequestSize.WithLabelValues(method, path).Observe(float64(c.Request.ContentLength))
		}

		// 处理请求
		c.Next()

		// 计算请求时长
		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Writer.Status())

		// 记录指标
		httpRequestsTotal.WithLabelValues(method, path, status).Inc()
		httpRequestDuration.WithLabelValues(method, path, status).Observe(duration)
		httpResponseSize.WithLabelValues(method, path).Observe(float64(c.Writer.Size()))
	}
}
