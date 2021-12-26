package models

import (
	"app/internal/dao"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo/field"
)

type Logger struct {
	field.DefaultField `bson:",inline"`
	UserId             primitive.ObjectID `bson:"userId,omitempty"`
	StartTime          time.Time          `bson:"startTime"`
	EndTime            time.Time          `bson:"endTime"`
	UseTime            string             `bson:"useTime"`
	IP                 string             `bson:"ip"`
	Method             string             `bson:"method"`
	Url                string             `bson:"url"`
	StatusCode         int                `bson:"statusCode"`
	Type               string             `bson:"type"`
	ErrMsg             string             `bson:"errMsg"`
	StackInfo          string             `bson:"stackInfo"`
}

func (logger *Logger) Start() {
	logger.StartTime = time.Now()
}

func (logger *Logger) End(c *gin.Context) {

	// 结束时间
	logger.EndTime = time.Now()

	// 执行时间
	logger.UseTime = logger.EndTime.Sub(logger.StartTime).String()

	// 请求方式
	logger.Method = c.Request.Method

	// 请求路由
	logger.Url = c.Request.RequestURI

	// 状态码
	logger.StatusCode = c.Writer.Status()

	// 请求IP
	logger.IP = c.ClientIP()

	if logger.ErrMsg != "" {
		logger.Type = "error"
	} else {
		logger.Type = "normal"
	}

	if err := dao.Insert("requestLogs", logger); err != nil {
		fmt.Println("日志记录失败")
	}
}
