package service

import (
	"k8s-platform/dao"
	"k8s-platform/model"
)

/**
 * @author       weimenghua
 * @time         2023-10-28 14:40
 * @description
 */

type workflow struct{}

var Workflow workflow

//获取列表分页查询
func (w *workflow) GetList(namespace string, page, limit int) (data *dao.WorkflowResp, err error) {
	data, err = dao.Workflow.GetList(namespace, page, limit)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//查询workflow单条数据
func (w *workflow) GetById(id int) (data *model.Workflow, err error) {
	data, err = dao.Workflow.GetById(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//定义WorkflowCreate结构体，用于创建workflow需要的参数属性的定义
type WorkflowCreate struct {
	Name          string                 `json:"name"`
	Namespace     string                 `json:"namespace"`
	Replicas      int32                  `json:"replicas"`
	Image         string                 `json:"image"`
	Label         map[string]string      `json:"label"`
	Cpu           string                 `json:"cpu"`
	Memory        string                 `json:"memory"`
	ContainerPort int32                  `json:"container_port"`
	HealthCheck   bool                   `json:"health_check"`
	HealthPath    string                 `json:"health_path"`
	Type          string                 `json:"type"`
	Port          int32                  `json:"port"`
	NodePort      int32                  `json:"node_port"`
	Hosts         map[string][]*HttpPath `json:"hosts"`
}

//workflow名字转换成ingress名字，添加-ing后缀
func getIngressName(workflowName string) (ingressName string) {
	return workflowName + "-ing"
}

//workflow名字转换成service名字，添加-svc后缀
func getServiceName(workflowName string) (ingressName string) {
	return workflowName + "-svc"
}
