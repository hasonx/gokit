package kits

import (
	"time"

	"github.com/hasonx/gokit/define"
)

// 将字符串格式的时间转换为int64
func StringToTimestamp(datetime string) int64 {
	theTime, _ := time.ParseInLocation(define.TimeFormat, datetime, time.Local)
	sr := theTime.Unix()
	return sr
}
