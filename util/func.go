package util

import (
	"Gin-Basic-Service/global"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"time"
)

func IndexOf(s string, sep string) int {

	for i := range s {
		if i+len(sep) > len(s) {
			break
		}
		str := string(s[i : i+len(sep)])
		if str == sep {
			return i
		}
	}

	return -1
}

func GetIPv4() (ip string) {
	// 获取所有网络接口的地址信息
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("无法获取网络接口地址:", err)
		return
	}

	// 遍历地址信息，找到与服务相关的IP地址
	for _, addr := range addrs {
		// 检查地址类型是否为IP地址
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			// 检查是否为IPv4地址（可以根据需要修改为IPv6）
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
			}
		}
	}
	return
}

type Record interface {
	SetRecord(userNumber string, time time.Time, action string)
}

func SetRecordFun(c *gin.Context, action string, record Record) error {
	userObj, bl := c.Get("user")
	if !bl {
		return errors.New("获取操作员失败！")
	}
	user := userObj.(*global.Operator)
	record.SetRecord(user.CUserNumber, time.Now(), action)

	return nil
}
