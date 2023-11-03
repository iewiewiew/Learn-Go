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



## 一、目录结构

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



## 二、详细步骤

### 2.1 后端

**2.1.1 前置准备**

创建目录  
`mkdir {config, controller, dao, db, midile, model, service, utils}`

批量下载依赖 注：在此项目行不通，需使用 `go get` 下载单个依赖  
```
go list -f '{{ join .Imports "\n" }}' ./... | grep -v "$(go list -f '{{ join .TestImports "\n" }}' ./...)" | grep -v "$(go list -f '{{ join .XTestImports "\n" }}' ./...)" | sort -u | xargs go get -v
```

下载单个依赖  
`go get <依赖名称>`

初始化项目  
`go mod init k8s-platform`

**2.2.2 项目开发**

1. 定义常量 config/config.go
2. 定义 service/init.go
3. 定义 main.go
4. 定义 service/dataselector.go
5. 定义 service/namespace.go
6. 定义 controller/namespace.go
7. 定义 controller/router.go
8. 定义 db.go
9. 定义 model/workflow.go
10. 定义 dao/workflow.go
11. 定义 service/workflow.go
12. 定义 controller/workflow.go
13. 定义 utils/jwt.go
14. 定义 middle/jwt.go
15. 在 mian.go 进行调用

**2.1.3 本地启动服务**  
kill -9 `lsof -i:9093 |awk '{print $2}' |tail -1` & go run main.go  

**2.1.4 访问接口**
- http://127.0.0.1:9093/api/k8s/namespaces
- http://127.0.0.1:9093/api/k8s/workflows?filter_name=&namespace=default&page=1&limit=10

**2.1.5 常用方法**

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



### 2.2 前端

**2.2.1 前置准备**

安装 vue  
`npm install -g @vue/cli`

初始化项目  
`vue init webpack k8s-platform-fe`

运行项目  
`cd k8s-platform-fe`
`npm run dev`

访问地址  
http://localhost:8081/#/

安装依赖
```
npm install element-ui -S
npm install element-plus --legacy-peer-deps
npm install nprogress
npm install jsonwebtoken
npm install @vue/shared --legacy-peer-deps
```

**2.2.2 项目开发**

1. 添加 vue.config.js
2. 修改 main.js
3. 添加 src/utils/request.js
4. 添加 assets
5. 添加 src/views/
6. 修改 src/router/index.js



## 三、项目部署
```
安装适用于Linux x86架构的交叉编译工具链
brew install FiloSottile/musl-cross/musl-cross

进入的Go项目的根目录，设置环境变量
export GOOS=linux
export GOARCH=386

构建项目
go build -o build/main main.go

将编译好的二进制文件复制到服务器上
scp build/main root@127.0.0.1:/usr/local/bin/

安装Go环境
yum -y install golang

在服务器上启动服务
nohup /usr/local/bin/main &
```
