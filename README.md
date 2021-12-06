# router-public-ip

获取本机公网ip，并将其发送到邮箱

本项目是通过淘宝的api获得本机的公网ip，并发送到自己的邮箱中。

## 使用

1. 克隆此项目

```shell
git@github.com:jimyag/router-public-ip.git
```

2. 安装依赖

```shell
go get
```

3. 修改`source/`中的接受者和发送者信息
4. 运行

```shell
go run main.go
```
