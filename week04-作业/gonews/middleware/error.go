package middleware

import (
	"app/internal/models"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

// 异常捕获
func CustomError(c *gin.Context) {

	defer func() {

		if err := recover(); err != nil {

			if c.IsAborted() {
				c.Status(200)
			}

			// 记录错误日志
			logger := &models.Logger{}
			logger.Start()

			// 调用栈信息
			buf := make([]byte, 2048)
			n := runtime.Stack(buf, false)
			stackInfo := fmt.Sprintf("%s", buf[:n])
			logger.StackInfo = stackInfo

			switch errStr := err.(type) {
			case string:
				logger.ErrMsg = errStr
				logger.End(c)
				c.JSON(200, gin.H{"code": 1, "data": nil, "msg": errStr})
			case error:
				logger.ErrMsg = errStr.Error()
				logger.End(c)
				fmt.Println(errStr)
				c.JSON(200, gin.H{"code": 1, "data": nil, "msg": "系统错误" + errStr.Error()})
			default:
				c.JSON(200, gin.H{"code": 1, "data": nil, "msg": "发生未知异常"})
			}
		}

	}()

	c.Next()

}
