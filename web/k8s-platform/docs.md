[TOC]

> By：weimenghua  
> Date：2023.10.27  
> Description：  

<h1 align="center">参考 k8s-platform 平台，抄一遍</h1>

**参考资料**
- [k8s-platform](https://gitee.com/wangzilong9570/k8s-platform)
- [k8s-platform-fe](https://gitee.com/wangzilong9570/k8s-platform-fe)
- [K8s管理系统项目实战【API开发】](https://www.yuque.com/wangzilong-4omf5/efv7oy/ssk1ix)
- [kubernetes 客户端库文档](https://kubernetes.io/zh-cn/docs/reference/using-api/client-libraries/)



### 目录结构
- config：定义全局配置，如监听地址、管理员账号等。
- controller：controller层，定义路由规则，及接口入参和响应。
- service：服务层，处理接口的业务逻辑。
- dao：数据库操作，包含数据库的增删改查。
- model：定义数据库的表的字段。
- db: 用于初始化数据库连接以及配置。
- middle：中间件层，添加全局的逻辑处理，如跨域、jwt验证等。
- utils：工具目录，定义常用工具，如token解析，文件操作等。
- go.mod：定义项目的依赖包以及版本。
- main.go：项目的主入口 main函数。



### 详细步骤
创建目录
mkdir {config, controller, dao, db, midile, model, servce, utils}

批量下载依赖 注：在此项目行不通，需使用go get下载单个依赖
go list -f '{{ join .Imports "\n" }}' ./... | grep -v "$(go list -f '{{ join .TestImports "\n" }}' ./...)" | grep -v "$(go list -f '{{ join .XTestImports "\n" }}' ./...)" | sort -u | xargs go get -v

下载单个依赖
go get <依赖名称>

初始化项目
go mod init k8s-platform

定义常量 config/config.go

定义 service/init.go

创建 main.go

定义 service/dataselector.go

定义 service/namespace.go

定义 controller/namespace.go

定义 controller/router.go

定义 db.go

定义 model/workflow.go

定义 dao/workflow.go

定义 service/workflow.go

定义 controller/workflow.go

本地启动服务
kill -9 `lsof -i:9093 |awk '{print $2}' |tail -1` & go run main.go
kill -9 `lsof -i:8080 |awk '{print $2}' |tail -1` & go run main.go

访问接口
http://127.0.0.1:9093/api/k8s/namespaces
http://127.0.0.1:9093/api/k8s/workflows?filter_name=&namespace=default&page=1&limit=10



### 常用方法
```
//获取pod列表
podList, err := K8s.ClientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
//获取pod详情
pod, err = K8s.ClientSet.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
//删除pod
err = K8s.ClientSet.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
//更新pod（完整yaml）
pod, err = K8s.ClientSet.CoreV1().Pods(namespace).Update(context.TODO(), pod, metav1.UpdateOptions{})
//获取deployment副本数
scale, err := K8s.ClientSet.AppsV1().Deployments(namespace).GetScale(context.TODO(), deploymentName, metav1.GetOptions{})
//创建deployment
deployment, err = K8s.ClientSet.AppsV1().Deployments(data.Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
//更新deployment（部分yaml）
deployment, err = K8s.ClientSet.AppsV1().Deployments(namespace).Patch(context.TODO(), deploymentName, "application/strategic-merge-patch+json", patchByte, metav1.PatchOptions{})
```



### 项目部署
```
yum -y install golang
```