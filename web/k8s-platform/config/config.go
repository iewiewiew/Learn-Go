package config

import "time"

/**
 * @author       weimenghua
 * @time         2023-10-27 20:25
 * @description
 */

const (
	//监听端口
	ListenAddr = "127.0.0.1:9093"
	//权限文件地址
	Kubeconfig = "/Users/menghuawei/.kube/config"
	//Kubeconfig = "~/.kube/config"
	//数据库配置
	DbType = "mysql"
	DbUser = "root"
	DbPwd  = "123456"
	DbHost = "127.0.0.1"
	DbPort = 3306
	DbName = "dbname"
	//连接池的配置
	MaxIdleConns = 10               //最大空闲连接
	MaxOpenConns = 100              //最大连接数
	MaxLifeTime  = 30 * time.Second //最大生存时间
	//日志显示行数
	PodLogTailLine = 2000
	//登录账户名和密码
	AdminUser = "admin"
	AdminPwd  = "admin"
)
