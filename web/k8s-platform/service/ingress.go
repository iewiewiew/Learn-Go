package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	nwv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/**
 * @author       weimenghua
 * @time         2023-10-31 09:59
 * @description
 */

type ingress struct{}

var Ingress ingress

//定义列表的返回内容，Items是ingress元素列表，Total为ingress元素数量
type IngressResp struct {
	Items []nwv1.Ingress `json:"items"`
	Total int            `json:"total"`
}

func (s *ingress) toCells(std []nwv1.Ingress) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = ingressCell(std[i])
	}
	return cells
}

func (s *ingress) fromCells(cells []DataCell) []nwv1.Ingress {
	ingresses := make([]nwv1.Ingress, len(cells))
	for i := range cells {
		ingresses[i] = nwv1.Ingress(cells[i].(ingressCell))
	}
	return ingresses
}

//定义ingress的path结构体
type HttpPath struct {
	Path         string `json:"path"`
	PathType     string `json:"path_type"`
	ServiceName  string `json:"service_name"`
	SevrvicePort int32  `json:"sevvice_port"`
}

//获取ingress列表，支持过滤、排序、分页
func (i *ingress) GetIngresses(filterName, namespace string, limit, page int) (ingressesResp *IngressResp, err error) {
	//获取ingressList类型的ingress列表
	ingressList, err := K8s.ClientSet.NetworkingV1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(errors.New("获取Ingress列表失败，" + err.Error()))
		return nil, errors.New("获取Ingress列表失败，" + err.Error())
	}
	//将ingressList中的ingress列表（Items），放进dataselector对象中，进行排序
	selectableData := &dataSelector{
		GenericDataList: i.toCells(ingressList.Items),
		dataSelectQuery: &DataSelectQuery{
			FilterQuery: &FilterQuery{
				Name: filterName,
			},
			PaginateQuery: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	data := filtered.Sort().Paginate()
	//将[]DataCell类型的ingress列表转为v1.ingress列表
	ingresses := i.fromCells(data.GenericDataList)
	logger.Info(ingresses)
	return &IngressResp{
		Items: ingresses,
		Total: total,
	}, nil
}
