package router

import (
	"Gin-Basic-Service/api/controller"
	"Gin-Basic-Service/router"
	"github.com/gin-gonic/gin"
)

type TableManageRouter struct {
	router *gin.RouterGroup
}

func init() {
	g := router.GetRouterGroup()
	r := TableManageRouter{}
	r.router = g.Group("/table_manage")
	r.appendRouter()
}

func (r *TableManageRouter) appendRouter() {
	con := controller.TableManageController{}
	r.router.POST("/create", con.CreateTable)
	r.router.PUT("/:id/update", con.UpdateTable)
	r.router.DELETE("/:id/delete", con.DeleteTable)
	r.router.GET("/:id", con.GetTableByID)
	r.router.POST("/query", con.QueryTable)
}
