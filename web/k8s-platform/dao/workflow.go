package dao

import (
	"errors"
	"github.com/wonderivan/logger"
	"k8s-platform/db"
	"k8s-platform/model"
)

/**
 * @author       weimenghua
 * @time         2023-10-28 13:40
 * @description
 */

type workflow struct{}

var Workflow workflow

// 定义列表的返回内容，Items是workflow元素列表，Total为workflow元素数量
type WorkflowResp struct {
	//[]*model.Workflow 表示一个切片，该切片可以存储多个指向 model.Workflow 类型对象的指针
	//[] 表示切片类型，表示可以存储多个元素的动态集合。
	//* 表示指针类型，它指向存储实际数据的内存地址，而不是直接存储数据本身。
	//model.Workflow 是指针切片中指针所指向的对象的类型。
	Items []*model.Workflow `json:"items"`
	Total int               `json:"total"`
}

// 获取列表分页查询
func (w *workflow) GetList(filterName, namespace string, page, limit int) (data *WorkflowResp, err error) {
	//定义分页数据的起始位置
	startSet := (page - 1) * limit
	//定义数据库查询返回内容
	var workflowList []*model.Workflow
	//数据库查询，Limit方法用于限制条数，Offset方法设置起始位置
	//tx := db.GORM.Where("namespace = ?", namespace). // @todo 暂时取消 namespace 查询条件
	tx := db.GORM.
		Where("name LIKE ?", "%"+filterName+"%").
		Limit(limit).
		Offset(startSet).
		Order("id desc").
		Find(&workflowList)
	//gorm会默认把空数据也放到err中，故这里要排除空数据的情况
	if tx.Error != nil && tx.Error.Error() != "record not found" {
		logger.Error("获取Workflow列表失败，" + tx.Error.Error())
		return nil, errors.New("获取Workflow列表失败，" + tx.Error.Error())
	}
	return &WorkflowResp{
		Items: workflowList,
		Total: len(workflowList),
	}, nil
}

// 获取workflow单条数据
func (w *workflow) GetById(id int) (workflow *model.Workflow, err error) {
	workflow = &model.Workflow{} //给空间
	tx := db.GORM.Where("id = ?", id).First(&workflow)
	if tx.Error != nil && tx.Error.Error() != "record not fuound" {
		logger.Error("获取Workflow单条数据失败，" + tx.Error.Error())
		return nil, errors.New("获取Workflow单条数据失败，" + tx.Error.Error())
	}
	return
}

// 新增workflow
func (w *workflow) Add(workflow *model.Workflow) (err error) {
	tx := db.GORM.Create(&workflow)
	if tx.Error != nil {
		logger.Error("添加workflow失败，" + tx.Error.Error())
		return errors.New("添加workflow失败，" + tx.Error.Error())
	}
	return
}

// 删除workflow
// 软删除db.GORM.Delete("id = ?",id)
// 软删除执行的是UPDATE语句，将deleted_at字段设置为时间即可，gorm默认就是软删
// 实际执行语句UPDATE 'workflow' SET 'deleted_at' = '2022-09-28 16:22:55' WHERE 'id' IN ('1')
// 硬删除db.GORM.Unscoped().Delete("id = ?",id)直接从表中删除这条数据
// 实际执行语句DELETE FROM 'workflow' WHERE 'id' IN ('1');
func (w *workflow) DelById(id int) (err error) {
	tx := db.GORM.Where("id = ?", id).Delete(&model.Workflow{})
	if tx.Error != nil {
		logger.Error("删除Workflow，" + tx.Error.Error())
		return errors.New("删除workflow失败，" + tx.Error.Error())
	}
	return nil
}