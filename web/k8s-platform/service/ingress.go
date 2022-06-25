package service

/**
 * @author       weimenghua
 * @time         2023-10-31 09:59
 * @description
 */

//定义ingress的path结构体
type HttpPath struct {
	Path         string `json:"path"`
	PathType     string `json:"pathType"`
	ServiceName  string `json:"serviceName"`
	SevrvicePort int32  `json:"sevrvicePort"`
}
