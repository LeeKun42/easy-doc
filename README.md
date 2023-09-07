# 微服务 用户模块

### 介绍
golang web项目
```
iris        框架
gorm        数据库
logrus      日志库
viper       配置文件读取
go-redis    redis库
```

### 项目结构
```
|-- app             //应用程序目录
|  |-- controller   //控制器文件目录
|  |-- lib          //自定义类库
|  |-- model        //数据模型定义目录
|  |-- service      //逻辑类目录
|-- config          //配置文件目录                 
|-- go.mod          //go mod依赖配置文件
|-- main.go         //主程序入口文件
|-- README.md
```


### 调试运行
```
    go run main.go 或者 air（热编译）
```

### 制作docker镜像
```
    编译镜像： docker build -t tagName . --no-cache
    运行镜像： docker run -it --rm --name=容器名称 -p 818:8108 tagName
    推送镜像： docker tag tagName:v1.0.0 127.0.0.1:5000/tagName:v1.0.0
              docker push 127.0.0.1:5000/tagName:v1.0.0
```
