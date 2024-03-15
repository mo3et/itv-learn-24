# Gin example 
参考煎鱼的实践文章
https://eddycjy.com/tags/gin/

> [Gin 1](https://eddycjy.com/posts/go/gin/2018-02-11-api-01/)

在前一章节中，我们初始化了一个 go-gin-example 项目，接下来我们需要继续新增如下目录结构：
```
go-gin-example/
├── conf
├── middleware
├── models
├── pkg
├── routers
└── runtime
```

- conf：用于存储配置文件
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
- routers 路由逻辑处理
- runtime：应用运行时数据
