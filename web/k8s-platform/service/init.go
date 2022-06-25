package service

import (
	"crypto/tls"
	"github.com/wonderivan/logger"
	"k8s-platform/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"time"
)

/**
 * @author       weimenghua
 * @time         2023-10-27 20:32
 * @description
 */

// k8s 结构体具有一个名为 ClientSet 的成员变量，其类型为 *kubernetes.Clientset。这意味着 ClientSet 是一个指向 kubernetes.Clientset 类型的指针。
type k8s struct {
	ClientSet *kubernetes.Clientset
}

// 定义了一个名为 k8s 的结构体类型，并声明了一个名为 K8s 的变量，其类型为 k8s 结构体。
var K8s k8s

func (k *k8s) Init() {
	// clientcmd.BuildConfigFromFlags 函数用于构建一个包含 Kubernetes 集群配置的 rest.Config 对象。
	conf, err := clientcmd.BuildConfigFromFlags("", config.Kubeconfig)
	if err != nil {
		panic("创建k8s配置失败，" + err.Error())
	}

	// 创建自定义的 http.Transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 跳过证书校验
		},
	}

	// 创建 http.Client，并设置自定义的 Transport
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}

	// 将自定义的 http.Client 设置到 Kubernetes 客户端配置中
	conf.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
		return client.Transport
	}

	//kubernetes.NewForConfig 函数用于根据给定的 rest.Config 对象创建一个 kubernetes.Clientset 客户端。
	clientSet, err := kubernetes.NewForConfig(conf)
	if err != nil {
		panic("创建K8s clientSet失败，" + err.Error())
	} else {
		logger.Info("创建k8s clientSet成功")
	}
	k.ClientSet = clientSet
}
