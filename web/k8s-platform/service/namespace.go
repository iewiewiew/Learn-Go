package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/**
 * @author       weimenghua
 * @time         2023-10-27 21:26
 * @description
 */

var Namespace namespace

type namespace struct{}

type NamespaceResp struct {
	Items []corev1.Namespace `json:"items"`
	Total int                `json:"total"`
}

//toCells方法用于将namespace类型数组，转换成DataCell类型数组
func (n *namespace) toCells(std []corev1.Namespace) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = namespaceCell(std[i])
	}
	return cells
}

//fromCells方法用于将DataCell类型数组，转换成namespace类型数组
func (n *namespace) fromCells(cells []DataCell) []corev1.Namespace {
	namespaces := make([]corev1.Namespace, len(cells))
	for i := range cells {
		namespaces[i] = corev1.Namespace(cells[i].(namespaceCell))
	}
	return namespaces
}

//获取namespace列表，支持过滤、排序、分页
func (n *namespace) GetNamespaces(filterName string, limit, page int) (namespaceResp *NamespaceResp, err error) {
	//获取namespaceList类型的namespace列表
	namespaceList, err := K8s.ClientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(errors.New("获取Namespace列表失败," + err.Error()))
		return nil, errors.New("获取Namespace列表失败," + err.Error())
	}
	//将namspaceList中的namespace列表(Items)，放进dataselector对象中，进行排序
	selecttableData := &dataSelector{
		GenericDataList: n.toCells(namespaceList.Items),
		dataSelectQuery: &DataSelectQuery{
			FilterQuery: &FilterQuery{Name: filterName},
			PaginateQuery: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	fileterd := selecttableData.Filter()
	total := len(fileterd.GenericDataList)
	data := fileterd.Sort().Paginate()
	//将[]DataCell类型的namespace列表转为v1.namespace列表
	namespaces := n.fromCells(data.GenericDataList)
	return &NamespaceResp{
		Items: namespaces,
		Total: total,
	}, nil
}

//获取namespace详情
func (n *namespace) GetNamespaceDeatil(namespaceName string) (namespace *corev1.Namespace, err error) {
	namespace, err = K8s.ClientSet.CoreV1().Namespaces().Get(context.TODO(), namespaceName, metav1.GetOptions{})
	if err != nil {
		logger.Error(errors.New("获取Namespace详情失败，" + err.Error()))
		return nil, errors.New("获取Namespace详情失败，" + err.Error())
	}
	return namespace, nil
}

//删除namespace
func (n *namespace) DeleteNamespace(namespaceName string) (err error) {
	err = K8s.ClientSet.CoreV1().Namespaces().Delete(context.TODO(), namespaceName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error("删除Namespace失败，" + err.Error())
		return errors.New("删除Namespace失败，" + err.Error())
	}
	return nil
}
