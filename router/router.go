package router

import (
	_ "Gin-Basic-Service/docs"
	"Gin-Basic-Service/global"
	"Gin-Basic-Service/middleware"
	"Gin-Basic-Service/util"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"sync"
)

func init() {
	// 设置 GIN_MODE 环境变量
	gin.SetMode(global.Conf.Server.Mode)
	global.Router = gin.Default()
	global.Router.Use(middleware.AuthMiddleware())
}

func GetRouterGroup() *gin.RouterGroup {
	g := global.Router.RouterGroup.Group(global.RouterPrefix)
	return g
}

func Run() {
	// 添加路由
	addRouter()
	// 创建一个 WaitGroup，用于等待服务器启动完成
	var wg sync.WaitGroup
	wg.Add(1)
	// 启动服务器并监听启动完成事件
	go func() {
		defer wg.Done()
		if err := global.Router.Run(fmt.Sprint(":", global.Conf.Server.Port)); err != nil {
			fmt.Println("启动服务器失败:", err)
		}
	}()
	// 在此处可以执行其他初始化操作或等待其他事件
	ip := util.GetIPv4()
	path := fmt.Sprint("http://", ip, ":", global.Conf.Server.Port)
	fmt.Println(fmt.Sprint("Api Path:     ", path))
	fmt.Println(fmt.Sprint("Swagger Path: ", path, "/swagger/index.html"))
	// 等待服务器启动完成
	wg.Wait()
}

func addRouter() {
	g := global.Router
	// 添加swagger路由
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 添加登录路由
	g.POST("/login", middleware.Login)
}
