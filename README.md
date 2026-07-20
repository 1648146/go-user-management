# go-user-management
基于 Gin+GORM 实现的用户信息管理后台系统
# Go 用户信息管理系统
## 项目介绍
基于Gin + GORM + MySQL开发的轻量化用户后台管理系统，实现用户账号注册、登录鉴权、账号状态管理、分页查询等基础后台能力。
适合Go后端入门练手，遵循标准MVC分层架构，代码结构清晰，可在此基础上扩展RBAC权限体系。

## 技术栈
- 编程语言：Go 1.21
- Web框架：Gin
- ORM框架：GORM
- 数据库：MySQL 8.0
- 附加：MD5密码加密、JWT登录令牌扩展支持

## 目录结构说明
| 文件夹     | 作用说明                     |
|------------|------------------------------|
| config     | 全局配置文件                 |
| model      | 数据库表映射结构体           |
| dao        | 底层数据库操作层              |
| service    | 业务逻辑处理层                |
| controller | 请求接收与响应处理层          |

## 本地启动教程
1. 创建MySQL数据库`go_user_system`
2. 修改config/app.yaml内MySQL连接地址与账号密码
3. 安装依赖
```shell
go mod tidy