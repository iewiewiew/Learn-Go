package service

import (
	"k8s-platform/dao"
	"k8s-platform/model"
)

/**
 * @author       weimenghua
 * @time         2024-07-20 17:23
 * @description
 */

type example struct {
}

var Example example

func (e *example) GetList(name string, page, limit int) (data *dao.ExampleResp, err error) {
	data, err = dao.Example.GetList(name, page, limit)
	return data, nil
}

type ExampleCreate struct {
	Name string `json:"name"`
	Num  int    `json:"num"`
}

func (e *example) CreateExample(data *ExampleCreate) (err error) {
	example := &model.Example{
		Name: data.Name,
		Num:  data.Num,
	}

	err = dao.Example.Add(example)
	if err != nil {
		return err
	}
	return nil
}

func (e *example) DelById(id int) (err error) {
	err = dao.Example.DelById(id)
	if err != nil {
		return err
	}
	return nil
}
