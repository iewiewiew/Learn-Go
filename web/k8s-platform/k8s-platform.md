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



## 1. 目录结构

- config：定义全局配置，如监听地址、管理员账号等。
- controller：controller 层，定义路由规则，及接口入参和响应。
- service：服务层，处理接口的业务逻辑。
- dao：数据库操作，包含数据库的增删改查。
- model：定义数据库的表的字段。
- db: 用于初始化数据库连接以及配置。
- middle：中间件层，添加全局的逻辑处理，如跨域、jwt 验证等。
- utils：工具目录，定义常用工具，如 token 解析，文件操作等。
- go.mod：定义项目的依赖包以及版本。
- main.go：项目的主入口 main 函数。



## 2. 详细步骤

### 2.1 后端

**2.1.1 前置准备**

创建目录  
```
mkdir {config, controller, dao, db, midile, model, service, utils}
```

批量下载依赖 注：在此项目行不通，需使用 `go get` 下载单个依赖  
```
go list -f '{{ join .Imports "\n" }}' ./... | grep -v "$(go list -f '{{ join .TestImports "\n" }}' ./...)" | grep -v "$(go list -f '{{ join .XTestImports "\n" }}' ./...)" | sort -u | xargs go get -v
```

下载单个依赖  
```
go get <依赖名称>
```

初始化项目  
```
go mod init k8s-platform
```

**2.2.2 项目开发**

1. 定义 config/config.go
2. 定义 service/dataselector.go
3. 定义 service/namespace.go
4. 定义 controller/namespace.go
5. 定义 controller/router.go
6. 定义 service/init.go
7. 定义 main.go
8. 定义 db.go
9. 定义 model/workflow.go
10. 定义 dao/workflow.go
11. 定义 service/workflow.go
12. 定义 controller/workflow.go
13. 定义 utils/jwt.go
14. 定义 middle/jwt.go
15. 在 mian.go 进行调用

**2.1.3 本地启动服务**  

```
kill -9 `lsof -i:9093 |awk '{print $2}' |tail -1` & go run main.go
```
注：本地测试接口时可禁用 main.go 里的 r.Use(middle.JWTAuth())

调试代码 注：未实践成功
```
go get -u github.com/go-delve/delve/cmd/dlv
dlv version
dlv debug

go build -gcflags "-N -l" main.go
gdb main
```

**2.1.4 访问接口**

- http://127.0.0.1:9093/api/k8s/namespaces
- http://127.0.0.1:9093/api/k8s/pods?filter_name=&namespace=default&page=1&limit=10
- http://127.0.0.1:9093/api/k8s/workflows?filter_name=&namespace=default&page=1&limit=10
- http://127.0.0.1:9093/api/k8s/examples?filter_name=&name=&page=1&limit=10

测试接口
```
查看 Example 列表
curl http://127.0.0.1:9093/api/k8s/examples\?filter_name=\&name=\&page=1\&limit=10

创建 Example
curl -X POST -H "Content-Type: application/json" http://127.0.0.1:9093/api/k8s/example/create -d '{"name":"hello1","num":111}'

删除 Example
curl -X DELETE -H "Content-Type: application/json" http://127.0.0.1:9093/api/k8s/example/del -d '{"id":1}'

查看 Pod 列表
curl http://127.0.0.1:9093/api/k8s/pods\?filter_name=\&namespace=default\&page=1\&limit=10
```

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

git submodule add git@gitee.com:wangzilong9570/k8s-platform-fe.git

修改 Config.js：const httpUrl='http://127.0.0.1:9093'; const wsUrl='ws://127.0.0.1:9091';
修复 Config.js：k8sNamespaceDetail: httpUrl+'/api/k8s/namespaces/detail', （namespace 改为 namespaces）

安装 vue  
```
npm install -g @vue/cli
```

初始化项目  
```
vue init webpack k8s-platform-fe
```

运行项目  
```
cd k8s-platform-fe
npm run dev
npm run serve
```

访问地址  
http://localhost:8081/#/

查看包列表
```
npm list
npm list --depth=0
```

安装依赖
```
npm install element-ui -S
npm install element-plus --legacy-peer-deps
npm install element-plus@2.6.0 --legacy-peer-deps
npm install nprogress
npm install jsonwebtoken
npm install @vue/shared --legacy-peer-deps
npm install --save @vue/reactivity --legacy-peer-deps
```

**2.2.2 项目开发**

1. 添加 vue.config.js
```
const { defineConfig } = require('@vue/cli-service')

module.exports = defineConfig({
  devServer: {
    host: '0.0.0.0', //监听地址
    port: 7070, //启动端口号
    open: true //启动后是否自动打开网页
  },
  transpileDependencies: true,
  lintOnSave: false //关闭语法检测
})
```

2. 修改 main.js
```
// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import { createApp } from 'vue'
//引入element plus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css' //引入图标视图
import * as ELIcons from '@element-plus/icons-vue'
//引入App.vue主题
import Vue from 'vue'
import App from './App'
//引入路由配置及规则
import router from './router'

// Vue.config.productionTip = false

/* eslint-disable no-new */
/*
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
*/

//创建vue实例
const app = createApp(App)

//将图标注册为全局组件
for (let iconName in ELTcons) {
  app.component(iconName, ELTcons[iconName])
}

//引入element plus
app.use(ElementPlus)

//引入路由
app.use(router)

//挂载
app.mount('#app')
```

3. 添加 src/utils/request.js
4. 添加 assets
5. 添加 src/views/
6. 修改 src/router/index.js



## 3. 项目部署

```
安装适用于 Linux x86 架构的交叉编译工具链
brew install FiloSottile/musl-cross/musl-cross

进入的 Go 项目的根目录，设置环境变量
export GOOS=linux
export GOARCH=386

构建项目
go build -o build/main main.go
go build -o /usr/local/bin/main /root/k8s-platform/main.go

将编译好的二进制文件复制到服务器上
scp build/main root@127.0.0.1:/usr/local/bin/

安装Go环境
yum -y install golang

在服务器上启动服务
nohup /usr/local/bin/main &

cd /usr/local/bin && go run main

docker build -t my-golang-app .

kubectl apply -f deployment.yaml
```



## 4. 知识碎片

初始化数据库

```
goose create create_example_table sql
goose -dir db mysql "root:123456@tcp(localhost:3306)/dbname?parseTime=true" up
goose -dir db mysql "root:123456@tcp(localhost:3306)/dbname?parseTime=true" up -f 20240720094036_create_example_table.sql
```

