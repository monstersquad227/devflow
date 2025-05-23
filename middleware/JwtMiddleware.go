package middleware

import (
	"devflow/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 需要放行的 API 白名单
var whiteList = map[string]bool{
	"/devflow/demo/testGet":    true,
	"/devflow/demo/testPost":   true,
	"/devflow/actuator/health": true,
	"/devflow/user/login":      true,
}

// Jwt 鉴权中间件
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求路径
		requestPath := c.Request.URL.Path

		// 如果在白名单中，则直接放行
		if _, exists := whiteList[requestPath]; exists {
			c.Next()
			return
		}

		tokenString := c.GetHeader("token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token required",
				"code":  3045})
			c.Abort()
			return
		}

		// 处理 Bearer 令牌格式
		//parts := strings.Split(tokenString, " ")
		//if len(parts) != 2 || parts[0] != "Bearer" {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		//	c.Abort()
		//	return
		//}

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  3046,
				"error": "无效的 token"})
			c.Abort()
			return
		}

		// 将用户名存入 context，后续处理可以使用
		c.Set("account", claims.Account)
		c.Next()
	}
}
