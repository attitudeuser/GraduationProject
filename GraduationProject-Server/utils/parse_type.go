package utils

import (
	"log"
	"strconv"
)

//类型转换相关函数在这里被定义

// ParseIntToString int 转 string
func ParseIntToString(num int) string {
	return strconv.Itoa(num)
}

// ParseInt64ToString int64 转 string
func ParseInt64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// ParseInt64ToInt int64 转 int
func ParseInt64ToInt(num int64) int {
	return int(num)
}

// ParseStringToInt string 转 int
func ParseStringToInt(s string) int {
	_t, e := strconv.Atoi(s)
	if e != nil {
		log.Fatal("parse string to int error:", e)
		return 0
	}
	return _t
}

// ParseStringToInt64 string 转 int64
func ParseStringToInt64(s string) int64 {
	_t, e := strconv.ParseInt(s, 10, 64)
	if e != nil {
		log.Fatal("parse string to int64 error:", e)
		return 0
	}
	return _t
}
