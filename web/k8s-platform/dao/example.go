package dao

import (
	"errors"
	"github.com/wonderivan/logger"
	"k8s-platform/db"
	"k8s-platform/model"
)

/**
 * @author       weimenghua
 * @time         2024-07-20 14:12
 * @description
 */

type example struct {
}

var Example example

type ExampleResp struct {
	Items []*model.Example `json:"items"`
	Total int              `json:"total"`
}

func (e *example) GetList(name string, page, limit int) (data *ExampleResp, err error) {
	startSet := (page - 1) * limit
	var exampleList []*model.Example
	//tx := db.GORM.Where("name = ?", name).
	tx := db.GORM.
		Limit(limit).
		Offset(startSet).
		Order("id desc").
		Find(&exampleList)
	if tx.Error != nil && tx.Error.Error() != "record not found" {
		logger.Error("获取 example 列表失败，" + tx.Error.Error())
		return nil, errors.New("获取 example 列表失败，" + tx.Error.Error())
	}
	return &ExampleResp{Items: exampleList, Total: len(exampleList)}, nil
}

func (e *example) Add(example *model.Example) (err error) {
	tx := db.GORM.Create(&example)
	if tx.Error != nil {
		logger.Error("添加 example 失败，" + tx.Error.Error())
		return errors.New(" 添加 example 失败，" + tx.Error.Error())
	}
	return
}

func (e *example) DelById(id int) (err error) {
	tx := db.GORM.Where("id = ?", id).Delete(&model.Example{})
	if tx.Error != nil {
		logger.Error("删除Example失败，" + tx.Error.Error())
		return errors.New("删除Example失败，" + tx.Error.Error())
	}
	return nil
}
