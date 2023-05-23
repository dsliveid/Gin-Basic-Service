package global

import (
	"time"
)

type Config struct {
	Server struct {
		Port            int    `yaml:"port"`
		Mode            string `yaml:"mode"`
		TokenExpiration string `yaml:"token_expiration"`
	} `yaml:"server"`
	Database struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"database"`
}

type Operator struct {
	CUserName   string `gorm:"column:c_user_name" json:"c_user_name"`     // 姓名
	CUserNumber string `gorm:"column:c_user_number" json:"c_user_number"` // 工号
}

// RecordInfo 操作信息公用结构体
type RecordInfo struct {
	CCreateBy   string     `gorm:"column:c_create_by" json:"c_create_by"`               // 创建人
	DCreateTime *time.Time `gorm:"column:d_create_time" json:"d_create_time"`           // 创建时间
	CUpdateBy   string     `gorm:"column:c_update_by" json:"c_update_by"`               // 修改人
	DUpdateTime *time.Time `gorm:"column:d_update_time" json:"d_update_time"`           // 修改时间
	CDeleteBy   string     `gorm:"column:c_delete_by" json:"c_delete_by,omitempty"`     // 删除人
	DDeleteTime *time.Time `gorm:"column:d_delete_time" json:"d_delete_time,omitempty"` // 删除时间
}
