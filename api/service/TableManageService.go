package service

import (
	"Gin-Basic-Service/api/model"
	"Gin-Basic-Service/global"
	"Gin-Basic-Service/util"
)

type TableManageService struct {
}

func (s *TableManageService) CreateTable(table *model.TableManage) error {
	db := global.DB
	err := db.Save(table).Omit("c_update_by", "d_update_time", "c_delete_by", "d_delete_time").Error
	return err
}

func (s *TableManageService) UpdateTable(table *model.TableManage) error {
	db := global.DB
	err := db.Save(table).Omit("c_create_by", "d_create_time", "c_delete_by", "d_delete_time").Error
	return err
}

func (s *TableManageService) DeleteTable(table *model.TableManage) error {
	db := global.DB
	err := db.Save(table).Omit("c_create_by", "d_create_time", "c_update_by", "d_update_time").Error
	return err
}

func (s *TableManageService) GetTableByID(id string) (*model.TableManage, error) {
	db := global.DB
	var table model.TableManage
	result := db.Model(&model.TableManage{})
	result.Where("id = ?", id)
	result.Where("d_delete_time is null")
	err := result.Preload("Fields", "d_delete_time is null").Find(&table).Error
	if err != nil {
		return nil, err
	}
	return &table, nil
}

func (s *TableManageService) QueryTable(p *util.PageInfo) (pageResult *util.PageResult, err error) {
	db := global.DB
	var obj *[]model.TableManage

	query := p.QueryInfo.(*model.TableManageQuery)
	result := db.Model(&model.TableManage{})
	if query.ID > 0 {
		result.Where("id=?", query.ID)
	}
	if query.CTable != "" {
		result.Where("c_table=?", query.CTable)
	}
	if query.CTableName != "" {
		result.Where("c_table like ?", query.CTableName+"%")
	}
	result.Where("d_delete_time is null")

	// 查询总行数
	var count int64
	count, err = p.GetCountDB(result)
	if err != nil {
		return
	}
	// 分页查询
	p.GetPageDB(result)
	err = result.Preload("Fields", "d_delete_time is null").Find(&obj).Error
	if err != nil {
		return
	}
	pageResult = &util.PageResult{}
	pageResult.Total = count
	pageResult.Items = obj
	return
}
