package model

import (
	"Gin-Basic-Service/global"
	"time"
)

// TableManage 表示(表信息表 ) table_manage 表的模型
type TableManage struct {
	ID            int64  `gorm:"primary_key;column:id" json:"id"`               // 编号
	CTable        string `gorm:"column:c_table" json:"c_table"`                 // 表名
	CTableName    string `gorm:"column:c_table_name" json:"c_table_name"`       // 表名称
	CTableDetails string `gorm:"column:c_table_details" json:"c_table_details"` // 表说明
	BEnable       int    `gorm:"column:b_enable" json:"b_enable"`               // 是否可用，0：不可用，1：可用
	global.RecordInfo
	Fields []FieldManage `gorm:"foreignkey:CTableID" json:"fields"` // 列信息
}

// TableManageAdd 新增的入参模型
type TableManageAdd struct {
	CTable        string           `gorm:"column:c_table" json:"c_table"`                 // 表名
	CTableName    string           `gorm:"column:c_table_name" json:"c_table_name"`       // 表名称
	CTableDetails string           `gorm:"column:c_table_details" json:"c_table_details"` // 表说明
	BEnable       int              `gorm:"column:b_enable" json:"b_enable"`               // 是否可用，0：不可用，1：可用
	Fields        []FieldManageAdd `json:"fields"`                                        // 列信息
}

// TableManageUpdate 修改的入参模型
type TableManageUpdate struct {
	ID int64 `gorm:"primary_key;column:id" json:"id"` // 编号
	TableManageAdd
	Fields []FieldManageUpdate `json:"fields"` // 列信息
}

// TableManageQuery 查询的入参模型
type TableManageQuery struct {
	ID         int64  `gorm:"primary_key;column:id" json:"id"`         // 编号
	CTable     string `gorm:"column:c_table" json:"c_table"`           // 表名
	CTableName string `gorm:"column:c_table_name" json:"c_table_name"` // 表名称
}

// TableName 返回模型对应的表名
func (t *TableManage) TableName() string {
	return "table_manage"
}

// SetRecord 设置操作人操作时间
func (t *TableManage) SetRecord(userNumber string, time time.Time, action string) {
	switch action {
	case "add":
		t.CCreateBy = userNumber
		t.DCreateTime = &time
	case "update":
		t.CUpdateBy = userNumber
		t.DUpdateTime = &time
	case "delete":
		t.CDeleteBy = userNumber
		t.DDeleteTime = &time
	}

}
