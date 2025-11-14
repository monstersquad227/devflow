package httpclient

import (
	"devflow/middleware"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPClient 带日志的 HTTP 客户端
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient 创建 HTTP 客户端
func NewHTTPClient(timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// DoRequest 执行 HTTP 请求并自动记录日志
func (c *HTTPClient) DoRequest(req *http.Request) ([]byte, int, error) {
	start := time.Now()

	// 执行请求
	resp, err := c.client.Do(req)
	latency := time.Since(start)

	if err != nil {
		// 记录请求失败日志
		middleware.LogHttpOut(req.Method, req.URL.String(), req.Header, req.Body, -9999, nil, latency, err)
		return nil, 0, fmt.Errorf("HTTP 请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		middleware.LogHttpOut(req.Method, req.URL.String(), req.Header, req.Body, resp.StatusCode, string(body), latency, err)
		return nil, resp.StatusCode, fmt.Errorf("读取响应失败: %w", err)
	}

	// 记录日志
	var logErr error
	if resp.StatusCode >= 400 {
		logErr = fmt.Errorf("HTTP 状态码错误: %d", resp.StatusCode)
	}
	middleware.LogHttpOut(req.Method, req.URL.String(), req.Header, req.Body, resp.StatusCode, string(body), latency, logErr)

	return body, resp.StatusCode, nil
}
