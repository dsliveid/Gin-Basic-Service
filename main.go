package main

import (
	_ "Gin-Basic-Service/api/router"
	"Gin-Basic-Service/global"
	"Gin-Basic-Service/router"
	"Gin-Basic-Service/util"
	"fmt"
)

// @title Swagger Gin-Basic-Service API
// @version 1.0
// @description 应用程序的简短描述.
// @schemes http
// @securityDefinitions.apikey token
// @name token
// @in header

func main() {
	Init()
	router.Run()

	ip := util.GetIPv4()
	fmt.Println("api:", ip, global.Conf.Server.Port)
	fmt.Println("swagger:", ip, global.Conf.Server.Port, "swagger/index.html")
}

func Init() {
	util.InitDB()
}
