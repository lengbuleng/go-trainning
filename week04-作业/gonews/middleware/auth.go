package middleware

import (
	"app/internal/models"
	"app/internal/service/user_service"

	"github.com/gin-gonic/gin"
)

// 身份验证
func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("token")
		if token == "" {
			c.JSON(401, gin.H{"code": 1, "msg": "token不能为空"})
			c.Abort()
			return
		}

		tokenSvc := &user_service.TokenSvc{}
		user, err := tokenSvc.ParseToken(token)
		if err != nil {
			c.JSON(401, gin.H{"code": 1, "msg": err.Error()})
			c.Abort()
			return
		}

		c.Set("userId", user.Id)
		c.Set("role", user.Role)

		// 记录请求日志
		logger := &models.Logger{}
		logger.Start()
		logger.UserId = user.Id

		c.Next()

		logger.End(c)

	}

}
