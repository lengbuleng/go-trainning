package moment

import (
	"strconv"
	"time"
)

// 获取今天日期
func GetToday(g bool) string {
	if g {
		return time.Now().Format("2006-01-02")
	}
	return time.Now().Format("20060102")
}

func GetTimeStampStr() string {
	timestamp := time.Now().Unix()
	return strconv.FormatInt(timestamp, 10)
}
