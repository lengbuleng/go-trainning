#### go gin web框架
##### 数据库为mongo,各层含义如下
1.api 加载路由提供外部接口  
2.cmd/main.go 项目入口，主要负责程序的启动、关闭、配置初始化等.
3.configs/config.yaml 放置配置文件模板, 或默认配置, 一般是 yaml 文件
4.internal/dao  一般都放置不希望作为第三包给他人使用的, 即是当作私有库.数据库驱动以及crud工具层  
5.internal/models 模型，数据持久化层  
6.internal/service 数据处理/逻辑层  
7.middleware gin中间件  
8.pkg 工具层   
9.test 额外的外部测试应用程序和测试数据。

config.yaml填写完配置  
执行: `go run main.go`
