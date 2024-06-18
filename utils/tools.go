package utils

import (
	"crypto/md5"
	"encoding/hex"
	"text/template"
	"time"
)

var TFunc = template.FuncMap{
	"UnixToTime": UnixToTime,
	"FormatTime": FormatTime,
}

// MD5加密
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// 时间戳转换成日期
func UnixToTime(timestamp int64) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转换成时间戳
func FormatTime(date string) int64 {
	t, _ := time.Parse("2006-01-02 15:04:05", date)
	return t.Unix()
}

func IsMobile(mobile string) bool {
	return true
}
