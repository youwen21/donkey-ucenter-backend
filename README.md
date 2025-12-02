# gofly framework

#### 介绍

gofly framework 是一款基于go语言的api开发框架

gofly-gen 工具根据mysql数据裤配置，自动生成 dal, dml, service, handler，和 router 代码

#### 软件架构

软件架构说明
本框架基于gin gorm开发

#### 软件分层

- dal 读取数据层
- dml 缓存 redis, lru 层
- service 业务逻辑层
- handler 入参接收，校验入参层
- middleware 接口认证校验， 接口内容加密解密

#### 安装教程

1. cd $GOPATH/src
2. git clone git@gitee.com:apiok/gofly_framework.git my-project
3. cd my-project && rm -rf .git
4. cp conf/config.toml.example conf/config.toml
5. go mod tidy
6. go run .

#### 使用说明

1. xxxx
2. xxxx
3. xxxx

#### 代码生成工具下载

> https://gitee.com/apiok/gofly-gen/releases/
> 下载相关平台执行文件， 放到系统$PATH中， 更名为:gofly-gen

#### 代码自动生成

代码生成依赖conf/config.toml中mysql_default配置
> gofly-gen apply

