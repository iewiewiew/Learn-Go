package service

import (
	corev1 "k8s.io/api/core/v1"
)

/**
 * @author       weimenghua
 * @time         2023-12-11 15:06
 * @description
 */

type configMap struct{}

var ConfigMap configMap

type ConfigMapsResp struct {
	Items []corev1.ConfigMap `json:"items"`
	Total int                `json:"total"`
}

func (c *configMap) toCells(std []corev1.ConfigMap) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = configMapCell(std[i])
	}
	return cells
}

func (c *configMap) fromCells(cells []DataCell) []corev1.ConfigMap {
	configMaps := make([]corev1.ConfigMap, len(cells))
	for i := range configMaps {
		configMaps[i] = corev1.ConfigMap(cells[i].(configMapCell))
	}
	return configMaps
}

//获取ConfigMap列表，支持过滤、分页、排序
