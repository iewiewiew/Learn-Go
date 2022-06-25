package service

import corev1 "k8s.io/api/core/v1"

/**
 * @author       weimenghua
 * @time         2023-10-31 10:33
 * @description
 */

//定义pod类型和Pod对象，用于包外的调用(包是指service目录)，例如Controller
var Pod pod

type pod struct{}

//toCells方法用于将pod类型的数组，转换成DataCell类型数组
func (p *pod) toCells(std []corev1.Pod) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = podCell(std[i])
	}
	return cells
}

//fromCells方法用于将DataCell类型数组，转换成pod类型数组
func (p *pod) fromCells(cells []DataCell) []corev1.Pod {
	pods := make([]corev1.Pod, len(cells))
	for i := range cells {
		//cells[i].(podCell)就使用到了断言，断言后转换成了podCell类型，然后又转换成了Pod类型
		pods[i] = corev1.Pod(cells[i].(podCell))
	}
	return pods
}

//定义列表返回内容，Item是pod元素列表，Total为pod元素数量
type PodsResp struct {
	Items []corev1.Pod `json:"items"`
	Total int          `json:"total"`
}

//定义PodsNp类型，用于返回namespace中的pod的数量
type PodsNp struct {
	Namespace string `json:"namespace"`
	PodNum    int    `json:"podNum"`
}
