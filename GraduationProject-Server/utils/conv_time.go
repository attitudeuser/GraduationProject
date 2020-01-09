package utils

import "time"

//返回当前日期时间
func GetNowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
