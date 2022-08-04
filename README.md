# Gin-Demo
## 简介
一个简单的Gin项目，依赖Gorm，Redis，Viper，Zap，JWT等库实现了一个基本的API服务器框架
## 项目基本架构

| 文件/目录          | 说明              |
|----------------|-----------------|
| app/common     | 公共模块(请求/响应结构体等) |
| app/controller | 业务调度器           |
| app/middleware | 中间件             |
| app/models     | 数据库结构体          |
| app/services   | 业务层             |
| bootstrap      | 项目初始化           |
| config         | 配置结构体           |
| global         | 全局变量            |
| routes         | 路由定义            |
| static         | 静态资源(允许外部访问)    |
| utils          | 工具函数            |
| config.yaml    | 配置文件            |
| main.go        | 项目启动文件          |

## 项目流程
1. 配置初始化，定义全局变量APP
2. 日志初始化，使用zap和lumberjack将日志写入文件中，加入全局变量Log
3. 数据库初始化，编写model并加入全局变量DB
4. 静态资源处理&优雅重启服务器，采用管道阻塞信号，并使用http.Server.Shutdown()方法关闭服务器
5. Validator(验证器)初始化，使用validator库，自定义验证器错误信息，自定义验证器，自定义错误码
6. 封装Response
7. 实现用户注册接口
8. 使用jwt-go包编写颁发Token逻辑，实现登录接口
9. 编写 jwt 鉴权中间件，避免在各个 controller 重复编写鉴权逻辑
10. 使用 jwt 鉴权中间件实现获取用户信息接口
11. 引入redis解决token注销问题(黑名单策略)
12. 封装分布式锁，在jwt中间件中增加续签机制
13. 使用文件记录Gin的错误日志，同时使用cors库做跨域处理

