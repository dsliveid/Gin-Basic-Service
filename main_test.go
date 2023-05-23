package main

import (
	"Gin-Basic-Service/api/model"
	"Gin-Basic-Service/util"
	"testing"
)

func TestName(t *testing.T) {
	queryInfo := &util.PageInfo{Page: 1, PageSize: 10}
	BindJSONToStruct(queryInfo, &model.TableManageQuery{})
}
func BindJSONToStruct(data interface{}, targetStruct interface{}) {
	// 解析 JSON 数据到 map[string]interface{} 对象
	dataMap := data.(string)
	println(dataMap)
}
