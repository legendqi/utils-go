/* coding: utf-8
@Time :   2021/3/4 下午4:52
@Author : legend
@File :   log_test.go
*/
package test

import (
	"testing"
	"time"
	"utils-go/logger"
	"utils-go/scrypt"
)

func TestLog(t *testing.T) {
	logger.SetLogConfig("info", []string{"file", "console"})
	n := 1
	for {
		<-time.After(time.Second)
		logger.Info("测试Info日志" + scrypt.GetRandomString(20))
		logger.Debug("测试Debug日志" + scrypt.GetRandomString(20))
		logger.Warn("测试Debug日志" + scrypt.GetRandomString(20))
		logger.Error("测试Error日志" + scrypt.GetRandomString(20))
		n++
		if n > 10 {
			break
		}
	}
	logger.SetLevel("debug")
	for {
		<-time.After(time.Second)
		logger.Info("测试Info日志" + scrypt.GetRandomString(20))
		logger.Debug("测试Debug日志" + scrypt.GetRandomString(20))
		logger.Warn("测试Debug日志" + scrypt.GetRandomString(20))
		logger.Error("测试Error日志" + scrypt.GetRandomString(20))
		n++
		if n > 20 {
			break
		}
	}
}