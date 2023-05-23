package model

import "Gin-Basic-Service/global"

// FieldManage 表示（字段信息表）field_manage 表的模型
type FieldManage struct {
	ID                int64  `gorm:"primary_key;column:id" json:"id"`                       // 主键
	CTableID          int64  `gorm:"column:c_table_id" json:"c_table_id"`                   // 表id，关联表：table_manage.id
	CField            string `gorm:"column:c_field" json:"c_field"`                         // 列名
	CFieldName        string `gorm:"column:c_field_name" json:"c_field_name"`               // 列名称
	CFieldDetails     string `gorm:"column:c_field_details" json:"c_field_details"`         // 列说明
	CFieldType        string `gorm:"column:c_field_type" json:"c_field_type"`               // 列类型
	IFieldLength      int    `gorm:"column:i_field_length" json:"i_field_length"`           // 字段长度
	BPrimaryKey       int    `gorm:"column:b_primary_key" json:"b_primary_key"`             // 是否为主键 0：非主键 1：主键
	CPrimaryRule      string `gorm:"column:c_primary_rule" json:"c_primary_rule"`           // 主键生成规则 input：手动输入 auto_increment：自增长 snowflake_id：雪花ID
	BForeignKey       int    `gorm:"column:b_foreign_key" json:"b_foreign_key"`             // 是否为外键 0：非外键 1：外键
	CAssociationTable int64  `gorm:"column:c_association_table" json:"c_association_table"` // 关联表
	CAssociationField int64  `gorm:"column:c_association_field" json:"c_association_field"` // 关联字段
	CDefaultValue     string `gorm:"column:c_default_value" json:"c_default_value"`         // 默认值
	BSynchronous      int    `gorm:"column:b_synchronous" json:"b_synchronous"`             // 是否已同步 0：未同步 1：已同步
	BEnable           int    `gorm:"column:b_enable" json:"b_enable"`                       // 是否可用
	global.RecordInfo
}

// FieldManageAdd 新增的入参模型
type FieldManageAdd struct {
	CTableID          int64  `gorm:"column:c_table_id" json:"c_table_id"`                   // 表id，关联表：table_manage.id
	CField            string `gorm:"column:c_field" json:"c_field"`                         // 列名
	CFieldName        string `gorm:"column:c_field_name" json:"c_field_name"`               // 列名称
	CFieldDetails     string `gorm:"column:c_field_details" json:"c_field_details"`         // 列说明
	CFieldType        string `gorm:"column:c_field_type" json:"c_field_type"`               // 列类型
	IFieldLength      int    `gorm:"column:i_field_length" json:"i_field_length"`           // 字段长度
	BPrimaryKey       int    `gorm:"column:b_primary_key" json:"b_primary_key"`             // 是否为主键 0：非主键 1：主键
	CPrimaryRule      string `gorm:"column:c_primary_rule" json:"c_primary_rule"`           // 主键生成规则 input：手动输入 auto_increment：自增长 snowflake_id：雪花ID
	BForeignKey       int    `gorm:"column:b_foreign_key" json:"b_foreign_key"`             // 是否为外键 0：非外键 1：外键
	CAssociationTable int64  `gorm:"column:c_association_table" json:"c_association_table"` // 关联表
	CAssociationField int64  `gorm:"column:c_association_field" json:"c_association_field"` // 关联字段
	CDefaultValue     string `gorm:"column:c_default_value" json:"c_default_value"`         // 默认值
	BSynchronous      int    `gorm:"column:b_synchronous" json:"b_synchronous"`             // 是否已同步 0：未同步 1：已同步
	BEnable           int    `gorm:"column:b_enable" json:"b_enable"`                       // 是否可用
}

// FieldManageUpdate 修改的入参模型
type FieldManageUpdate struct {
	ID int64 `gorm:"primary_key;column:id" json:"id"` // 编号
	FieldManageAdd
}

// TableName 返回模型对应的表名
func (FieldManage) TableName() string {
	return "field_manage"
}
