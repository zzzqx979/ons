# UQValley Ons

## 一、项目简介

>- 目的和背景
> 
>       优刻谷作为一个物联网科技公司，基于自研的RFID射频标签，计划打造一个物联平台实现万物互联的功能。
>       ONS服务就是物联平台的核心服务通过任何物体RFID标签的TID就能够查到对应的物模型，有了物模型就可以对其进行管理和控制。
>- 功能
>       
>       我们的ONS服务集成了ONS和MS服务，通过设备的TID可以查询到设备的物模型。ONS服务还分成两个版本，分别是云平台版本和网关版本。两个版本都实现了基本的查询功能，云平台版本对比网关版本增加了数据存储功能。
>       ONS全称Object Name Service，对象名称解析服务。作用类似于DNS，先将TID换取EPC数据，再通过EPC数据获取对应的模型URL。
>       MS：Model Service，模型解释服务。通过访问指定模型URL，获取其对应的物模型。(待完善)

## 二、环境依赖

> GOLANG V1.21.1 + MYSQL + REDIS + EMQX

## 三、配置说明

> [项目开发/构建/运行时的配置说明，如：环境变量、配置文件等。]

## 四、开发说明
- yaml文件
```yaml
    debug: true
    server:
      port: 8080
    db:
      type: mysql
      source: username:password@tcp(ip:port)/db?charset-utf8
    redis:
      addr: localhost:6379
      password:
      db:
    mqtt:
      broker: localhost:18083
      username: username
      password: password

```
- 部署命令
```shell
    // 使用yaml文件部署
    go run main.go aiot
```
### 部署网关服务
- yaml文件
```yaml
   debug: true
   server:
     port: 8081
   db:
     type: sqlite3
     source: ./output/gateway.db

```
- 部署命令
```shell
    // 使用yaml文件部署
    go run main.go gateway
```
- 服务设置
```shell
// 下面的操作需要先将程序的可执行文件和yaml配置文件放到/home/user/ons目录下,具体目录可以自行配置
// 配置服务启动文件
user@Phytium-Pi:~$ vim /lib/systemd/system/ons.service
[Unit]
Desciption=gateway ons server

[Service]
Type=simple
Restart=on-failure
User=user
ExecStart=/home/user/ons/ons gateway -c /home/user/ons/gateway1.yaml

[Install]
WantedBy=multi-user.target
// 设置ons服务开机自启动
user@Phytium-Pi:~$ sudo systemctl enable ons.service
// 开启ons服务
user@Phytium-Pi:~$ sudo systemctl start ons.service
```

## 五、更新记录

### v0.1.0.20240117

⭐ New Features

- 【新增】厂商管理，产品管理，tid管理，物模型管理；
- 【新增】ONS服务和MS服务；


## 六、常见问题

> [列出并回答与项目相关的常见问题。]

### 问题描述xxxxxxxx；

原因分析解决方案xxxxxxxx；
