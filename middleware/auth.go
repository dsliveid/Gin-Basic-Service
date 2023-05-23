package middleware

import (
	"Gin-Basic-Service/global"
	"Gin-Basic-Service/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// 定义全局变量，用于排除不需要token验证的接口
var excludedPaths = []string{"/login", "/logout", "/swagger/*"}

// AuthMiddleware 定义中间件函数，用于验证登录状态
func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 判断是否为需要token验证的接口，如果不是则退出
		if strings.HasPrefix(c.Request.URL.Path, global.RouterPrefix) {
			url := strings.Replace(c.Request.URL.Path, global.RouterPrefix, "", -1)
			for _, path := range excludedPaths {
				if strings.HasSuffix(path, "/*") {
					if strings.HasPrefix(url, path) {
						c.Next()
						return
					}
				} else {
					if c.Request.URL.Path == path {
						c.Next()
						return
					}
				}
			}
		} else {
			c.Next()
			return
		}

		// 从请求头中获取 Token
		token := c.GetHeader("token")
		// 判断 Token 是否为空
		if token == "" {
			util.Respond(c, global.Unauthorized, "Unauthorized 未获取到token!", nil)
			c.Abort() // 终止请求处理流程
			return
		}
		// Token 的解析和验证
		claims, err := util.ParseToken(token)
		if err != nil {
			util.Respond(c, global.Unauthorized, "Unauthorized token解析失败!", nil)
			c.Abort() // 终止请求处理流程
			return
		}

		// 验证通过，将用户信息存储到上下文中，供后续处理函数使用
		c.Set("user", claims.Operator)

		c.Next() // 继续执行后续的处理函数
	}
}

// Login godoc
// @Summary 登录接口
// @Description Login
// @Accept json
// @Produce json
// @Param table body global.Operator true "表信息"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	user := global.Operator{CUserName: "admin", CUserNumber: "001"}
	token, err := util.GenerateToken(&user)
	if err != nil {
		util.Respond(c, global.ErrorCode, fmt.Sprint("登录失败!", err.Error()), nil)
		return
	}
	// 假设登录成功
	util.Respond(c, global.SuccessCode, "登录成功!", gin.H{"Token": token})
}
