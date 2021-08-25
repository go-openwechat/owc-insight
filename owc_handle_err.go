package main

import (
	"time"
)

var count int32

// 错误处理函数
func messageErrorHandler(err error) {
	t := time.Now()
	if t.Sub(lastError) < 30*time.Minute {
		count++
	} else {
		count = 1
	}
	// 如果发生了三次错误,那么直接退出
	if count > 3 {
		abortOn("Too many errors", err)
	}
	logIf(0, "catch-and-skip", "count", count, "err", err)
	lastError = time.Now()
	chatie.SendText("ding")
	go wxHandshakeCheck()
}
