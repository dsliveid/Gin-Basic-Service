package controller

import (
	"Gin-Basic-Service/api/model"
	"Gin-Basic-Service/api/service"
	"Gin-Basic-Service/global"
	"Gin-Basic-Service/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TableManageController struct {
}

// CreateTable godoc
// @Summary 创建表
// @Description 创建新的表
// @Tags table_manage
// @Accept json
// @Produce json
// @Param table body model.TableManageAdd true "表信息"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security token
// @Router /api/custom/routing/table_manage/create [post]
func (c *TableManageController) CreateTable(ctx *gin.Context) {
	service := service.TableManageService{}
	table := &model.TableManage{}
	err := ctx.ShouldBindJSON(table)

	if err != nil {
		util.Respond(ctx, global.ErrorCode, fmt.Sprint("参数解析失败！", err.Error()), nil)
		return
	}

	util.SetRecordFun(ctx, "add", table)
	err = service.CreateTable(table)
	if err != nil {
		util.Respond(ctx, global.ErrorCode, fmt.Sprint("新增数据失败！", err.Error()), nil)
		return
	}
	util.Respond(ctx, global.SuccessCode, "新增成功！", table)
}

// UpdateTable godoc
// @Summary 更新表信息
// @Description 更新指定表的信息
// @Tags table_manage
// @Accept json
// @Produce json
// @Param table body model.TableManageUpdate true "表信息"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security token
// @Router /api/custom/routing/table_manage/{id}/update [put]
func (c *TableManageController) UpdateTable(ctx *gin.Context) {
	service := service.TableManageService{}
	table := &model.TableManage{}
	err := ctx.ShouldBindJSON(table)
	if err != nil {
		util.Respond(ctx, global.ErrorCode, fmt.Sprint("参数解析失败！", err.Error()), nil)
		return
	}

	util.SetRecordFun(ctx, "update", table)
	err = service.UpdateTable(table)
	if err != nil {
		util.Respond(ctx, global.ErrorCode, fmt.Sprint("修改失败！", err.Error()), nil)
		return
	}
	util.Respond(ctx, global.SuccessCode, "修改成功！", table)
}

// DeleteTable godoc
// @Summary 删除表
// @Description 删除指定表
// @Tags table_manage
// @Accept json
// @Produce json
// @Param id path string true "表ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security token
// @Router /api/custom/routing/table_manage/{id}/delete [delete]
func (c *TableManageController) DeleteTable(ctx *gin.Context) {
	service := service.TableManageService{}
	id := ctx.Param("id")
	table := &model.TableManage{}
	var err error
	table.ID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		util.Respond(ctx, global.ErrorCode, fmt.Sprint("入参：id类型转换失败！", err.Error()), nil)
		return
	}

	util.SetRecordFun(ctx, "delete", table)
	err = service.DeleteTable(table)
	if err != nil {
		util.Respond(ctx, global.ErrorCode, fmt.Sprint("删除失败！", err.Error()), nil)
		return
	}
	util.Respond(ctx, global.SuccessCode, "删除成功！", gin.H{"id": id})
}

// GetTableByID godoc
// @Summary 获取表信息
// @Description 获取指定表的信息
// @Tags table_manage
// @Accept json
// @Produce json
// @Param id path string true "表ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security token
// @Router /api/custom/routing/table_manage/{id} [get]
func (c *TableManageController) GetTableByID(ctx *gin.Context) {
	service := service.TableManageService{}
	id := ctx.Param("id")
	table, err := service.GetTableByID(id)
	if err != nil {
		util.Respond(ctx, global.ErrorCode, fmt.Sprint("查询失败！", err.Error()), nil)
		return
	}
	util.Respond(ctx, global.SuccessCode, "查询成功！", table)
}

// QueryTable godoc
// @Summary 获取表信息列表
// @Description 获取表的信息列表
// @Tags table_manage
// @Accept json
// @Produce json
// @Param table body util.PageInfo{query_info=model.TableManageQuery} true "表信息"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security token
// @Router /api/custom/routing/table_manage/query [post]
func (c *TableManageController) QueryTable(ctx *gin.Context) {
	service := service.TableManageService{}

	queryInfo := &util.PageInfo{QueryInfo: &model.TableManageQuery{}}
	err := ctx.ShouldBindJSON(queryInfo)
	if err != nil {
		util.Respond(ctx, global.ErrorCode, fmt.Sprint("参数解析失败！", err.Error()), nil)
		return
	}
	page, err := service.QueryTable(queryInfo)
	if err != nil {
		util.Respond(ctx, global.ErrorCode, fmt.Sprint("查询失败！", err.Error()), nil)
		return
	}
	util.Respond(ctx, global.SuccessCode, "查询成功！", page)
}
