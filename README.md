# yq-starter

基于gin和gorm构建web应用

- gin
- gorm
- viper
- vgo
- logrus

## vgo命令
```shell
vgo version 查看vgo版本
vgo install 安装依赖
vgo build 编译项目
vgo run 运行项目
vgo get github.com/gin-gonic/gin 获取依赖包的最新版本
vgo get github.com/gin-gonic/gin@v1.2 获取依赖包的指定版本
vgo mod -vendor 将依赖包直接放在项目的vendor目录里
```

## Start

```shell
vgo run app.go
```

## Build

```shell
vgo build app.go
```

## Deploy

```shell
nohup ./app &
```

## Todo

- redis
- kafka
- rpc
- mirc service