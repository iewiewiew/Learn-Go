package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

/**
 * @author       weimenghua
 * @time         2023-11-28 15:38
 * @description
 */

type service struct{}

var Service service

//定义列表返回内容，Items是deployment元素列表，total是deployment元素数量
type ServicesResp struct {
	Items []corev1.Service `json:"items"`
	Total int              `int:"total"`
}

//定义DeploymentNp类型，用于返回namespace中deployment数量
type ServiceNp struct {
	Namespace  string `json:"namespace"`
	ServiceNum int    `json:"service_num"`
}

func (s *service) toCells(std []corev1.Service) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = serviceCell(std[i])
	}
	return cells
}

func (s *service) fromCells(cells []DataCell) []corev1.Service {
	services := make([]corev1.Service, len(cells))
	for i := range cells {
		services[i] = corev1.Service(cells[i].(serviceCell))
	}
	return services
}

//定义ServiceCreate结构体，用于创建service需要的参数属性的定义
type ServiceCreate struct {
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Type          string            `json:"type"`
	ContainerPort int32             `json:"container_port"`
	Port          int32             `json:"port"`
	NodePort      int32             `json:"node_port"`
	Label         map[string]string `json:"label"`
}

//获取service列表，支持过滤、排序、分页
func (s *service) GetServices(filterName, namespace string, limit, page int) (servicesResp *ServicesResp, err error) {
	//获取serviceList类型的service列表
	serviceList, err := K8s.ClientSet.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(errors.New("获取Service列表失败," + err.Error()))
		return nil, errors.New("获取Service列表失败," + err.Error())
	}
	//将serviceList中的service列表（Item），放进dataselector对象中，进行排序
	selectableData := &dataSelector{
		GenericDataList: s.toCells(serviceList.Items),
		dataSelectQuery: &DataSelectQuery{
			FilterQuery: &FilterQuery{Name: filterName},
			PaginateQuery: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	data := filtered.Sort().Paginate()
	//将[]DataCell类型的service列表转为v1.service列表
	services := s.fromCells(data.GenericDataList)
	return &ServicesResp{
		Items: services,
		Total: total,
	}, nil
}

//获取service详情
func (s *service) GetServiceDeatil(serviceName, namespace string) (service *corev1.Service, err error) {
	service, err = K8s.ClientSet.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		logger.Error(errors.New("获取Service详情失败，" + err.Error()))
		return nil, errors.New("获取Service详情失败，" + err.Error())
	}
	return service, nil
}

//创建service，接收ServiceCreate对象
func (s *service) CreateService(data *ServiceCreate) (err error) {
	//将data中的数据组装成corev1.Service对象
	service := &corev1.Service{
		//ObjectMeta定义资源名、命名空间以及标签
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},
		//Spec中定义类型，端口，选择器
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType(data.Type),
			Ports: []corev1.ServicePort{
				{
					Name:     "http",
					Port:     data.Port,
					Protocol: "TCP",
					TargetPort: intstr.IntOrString{
						Type:   0,
						IntVal: data.ContainerPort,
					},
				},
			},
			Selector: data.Label,
		},
	}
	//默认ClusterIP，这里判断NodePort，添加配置
	if data.NodePort != 0 && data.Type == "NodePort" {
		service.Spec.Ports[0].NodePort = data.NodePort
	}
	//创建Service
	_, err = K8s.ClientSet.CoreV1().Services(data.Namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		logger.Error(errors.New("创建Service失败，" + err.Error()))
		return errors.New("创建Service失败，" + err.Error())
	}
	return nil
}

//删除service
func (s *service) DeleteService(serviceName, namespace string) (err error) {
	err = K8s.ClientSet.CoreV1().Services(namespace).Delete(context.TODO(), serviceName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(errors.New("删除Service失败，" + err.Error()))
		return errors.New("删除Service失败，" + err.Error())
	}
	return nil
}

//更新service
func (s *service) UpdateService(namespace, content string) (err error) {
	var service = &corev1.Service{}
	err = json.Unmarshal([]byte(content), service)
	if err != nil {
		logger.Error(errors.New("反序列化失败，" + err.Error()))
		return errors.New("反序列化失败，" + err.Error())
	}
	_, err = K8s.ClientSet.CoreV1().Services(namespace).Update(context.TODO(), service, metav1.UpdateOptions{})
	if err != nil {
		logger.Error(errors.New("更新Service失败，" + err.Error()))
		return errors.New("更新Service失败，" + err.Error())
	}
	return nil
}
