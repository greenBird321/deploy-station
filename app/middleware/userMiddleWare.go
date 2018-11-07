package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.Abort()
}

// 标准http接口 中间件
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Request.Cookie("uid")
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}
		c.Next()
		return
	}
}
