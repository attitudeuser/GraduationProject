# GraduationProject-Server

这是一个基于Clean-Architecture的Go框架实现，Clean Architecture是Robert C. Martin提出的一种业务领域分层的方法论，文章参考[https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)

说明:结构架构是我在日常使用时结合编码习惯实现的简洁模型，开源在[https://github.com/togo-soft/Togo](https://github.com/togo-soft/Togo)

**目录**
* [框架结构](#框架结构)
* [应用流程](#应用流程)
* [依赖](#依赖)
* [使用说明](#使用说明)
* [分层测试](#分层测试)
* [参考](#参考)

## 框架结构

1. models实体层

实体层的数据是系统数据类型的抽象,是内部的关键业务逻辑结构，最不可能发生变化的部分。在项目中，该层的结构可能被用作参数的**类型**，进行传递。

2. repositories存储层

存储层一般用来进行CRUD操作，通常DB连接句柄或者ORM在此处使用。**此处没有业务流程**而是专注于数据的操作(DB，RPC，FS)，**这里没有业务流程**，只有普通的数据操作功能。

repo层定义了DB的操作接口，实现repo层与上层的隔离

3. usecases用例层

用例层时业务流程的处理程序，相当于MVC中的Controller控制器，该层通过业务流程决定使用哪个存储库层，并提供数据。

usecases层定义了业务的操作接口，实现与上层的隔离。它会接收来自调用者的请求，通过逻辑处理，返回处理后的数据(usecases不考虑是谁调用了它，只需要返回数据即可)。

4. deliveries交付层

决定数据如何呈现。这一层将接收来自用户的输入， 并清理数据然后传递给接口适配层。

在程序中，我将路由router和请求处理handle放在此处，router专门负责路由的定义，handle专门负责usecase业务逻辑的调用与数据交付(以REST API的方式)。

5. utils包

某些功能可能抽离成单独的函数方法比较合适(比如配置信息的读取)，utils包正在做这些事情，它隔离于框架，为程序运行时提供额外的组态。

6. middleware中间件

拦截路由，进行中间环节处理。

```text
project                  应用部署目录
├─deliveries             交付层
│  ├─handler             路由请求处理目录
│  ├─router              路由配置目录
│  └─run.go              启动器
├─usecases               实例层
├─repositories           仓库存储层
├─models                 实体层
├─utils                  常用工具
├─main.go                程序启动入口
├─config.yaml            配置样例文件
├─LICENSE                授权说明文件
├─README.md              README 文件
```

## 应用流程

![APP.png](https://i.loli.net/2019/10/03/T8KlSREIzakHcpj.png)

## 依赖

1. repositories层:[xorm](https://xorm.io)

   xorm: eXtra ORM for Go

2. deliveries层:[gin](https://github.com/gin-gonic/gin)

   gin:  A HTTP web framework written in Go

## 使用说明

真机运行:

```shell
//修改 config.yaml 的 mysql_dsn 数据库配置项
//运行
go run .
```

Dockerfile:

```bash
//先启用mysql,此处使用docker方式运行
docker run -it -d --name bs  -p 33061:3306 -e MYSQL_ROOT_PASSWORD=11ysc1s1kR -e MYSQL_DATABASE=bs mysql:latest --character-set-server=utf8mb4 --collation-server=utf8mb4_general_ci
//修改 config.yaml 的 mysql_dsn

//构建镜像
docker build -t togo/bs:v1 .
//运行容器 PATH为项目目录 expose_port为本机曝露(转发)端口
docker run -d -v PATH:/Project -p expose_port:5000 --name togo_app togo/bs:v1
```

Docker-compose:

```bash
//构建
docker-compose build --no-cache
//启动
docker-compose start
//注意:请修改 config.yaml 的 mysql_dsn HOST为 "mysql_c" 字符串
```

## 分层测试

自测。

## 参考

1. The Clean Architecture [https://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html](https://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html)

2. Clean Architecture Example (Java) [https://github.com/mattia-battiston/clean-architecture-example](https://github.com/mattia-battiston/clean-architecture-example)

3. go-arch-clean [https://github.com/bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)

4. Trying Clean Architecture on Golang [https://hackernoon.com/golang-clean-archithecture-efd6d7c43047](https://hackernoon.com/golang-clean-archithecture-efd6d7c43047)

## License

© xuthus, 2019~time.Now

This software is licensed under the [MIT License](./LICENSE)
